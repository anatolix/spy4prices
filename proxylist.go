package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Proxy struct {
	URL    *url.URL
	client *http.Client
	OK     bool
}

type ProxyList struct {
	proxies  []Proxy
	curProxy int32
}

func (l *ProxyList) LoadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	errCount := 0
	exists := make(map[string]int)

	for scanner.Scan() {
		args := strings.Split(scanner.Text(), ",")
		if len(args) < 6 {
			errCount++
			continue
		}
		host := args[1]
		port, err := strconv.Atoi(args[2])
		proto := args[4]
		if err != nil || len(host) == 0 || (proto != "http" && proto != "socks5") {
			errCount++
			continue
		}

		rawurl := fmt.Sprintf("%s://%s:%d", proto, host, port)
		if _, found := exists[rawurl]; found {
			continue
		}

		exists[rawurl] = 1

		url, err := url.Parse(rawurl)
		if err == nil {
			l.proxies = append(l.proxies, Proxy{URL: url})
		} else {
			errCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	log.Printf("%d proxies loaded from %s (%d errors)", len(l.proxies), fileName, errCount)
	return nil
}

func (l *ProxyList) CheckProxyList(timeout time.Duration) {
	log.Printf("Checking proxies. Wait for %v  ...", timeout)
	wg := sync.WaitGroup{}
	wg.Add(len(l.proxies))
	for idx := range l.proxies {
		go func(idx int) {
			client := &http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyURL(l.proxies[idx].URL),
					// MaxIdleConns:        100,
					// MaxIdleConnsPerHost: 100,
					// IdleConnTimeout:     90 * time.Second,
					// DisableKeepAlives: true,
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
				Timeout: timeout,
			}
			defer wg.Done()

			if resp, err := client.Get("https://ya.ru"); err == nil {
				defer resp.Body.Close()

				if _, err := ioutil.ReadAll(resp.Body); err == nil {
					l.proxies[idx].client = client
					l.proxies[idx].OK = true
					// client.Timeout = 30 * time.Second
				} else if err != nil {
					// log.Println(err)
				}
			} else if err != nil {
				// log.Println(err)
			}
		}(idx)
	}
	wg.Wait()
	j := len(l.proxies)

	for i := 0; i < j; i++ {
		if !l.proxies[i].OK {
			j--
			l.proxies[i], l.proxies[j] = l.proxies[j], l.proxies[i]
			i--
		}
	}
	l.proxies = l.proxies[0:j]
	log.Printf("After checking left %d proxies", j)

}

func (l *ProxyList) Next() Proxy {
	cur := int(atomic.AddInt32(&l.curProxy, 1))

	return l.proxies[cur%len(l.proxies)]
}
