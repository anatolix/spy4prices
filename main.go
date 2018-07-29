package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var ymCategories = []int{54440, 54425, 54419, 54432, 54422, 54421, 54496, 54436, 54438, 54418, 54423, 54420, 54434, 54430}

const numWorkers = 2000
const httpTimeout = time.Second * 30

var proxyList ProxyList
var uaList UAList
var ymWorkers *YMWorkers
var rxRepo RxRepo

// Tool to update proxies list: https://github.com/chill117/proxy-lists

func main() {

	fmt.Println("Spy4prices :)")

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	if err := proxyList.LoadFromFile("data/proxies.csv"); err != nil {
		log.Fatal(err)
	}

	if err := rxRepo.Init("cproto://reindexer.org:6534/yamarket"); err != nil {
		log.Println(err)
	}

	if err := uaList.LoadFromFile("data/useragents.txt"); err != nil {
		log.Fatal(err)
	}
	if err := proxyList.LoadFromFile("data/proxies.csv"); err != nil {
		log.Fatal(err)
	}
	proxyList.CheckProxyList(httpTimeout)

	var sess *YMSession

	ymWorkers = NewYMWorkers(numWorkers)

	for _, cat := range ymCategories {
		for {
			sess = ymWorkers.PostJobs(cat, sess)
			if sess != nil {
				break
			}
		}
	}
	ymWorkers.Wait()
}
