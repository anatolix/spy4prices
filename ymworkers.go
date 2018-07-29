package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type Job struct {
	Nid, CatID, Hid []string
	Page            int
}

type YMWorkers struct {
	jobs        chan Job
	doneWorkers int32
	postedJobs  int
	wg          sync.WaitGroup
	lock        sync.RWMutex
	failedQ     map[string]int
}

func NewYMWorkers(numWorkers int) *YMWorkers {
	w := &YMWorkers{
		jobs:    make(chan Job, numWorkers*2),
		failedQ: make(map[string]int),
	}
	w.wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go w.run()
	}
	return w
}

func (w *YMWorkers) Wait() {
	close(w.jobs)
	w.wg.Wait()
}

func (w *YMWorkers) PostJobs(category int, sess *YMSession) *YMSession {

	if sess == nil {
		sess = NewYMSession(proxyList.Next().client, uaList.Next())

		if err := sess.Start(); err != nil {
			log.Println(err)
			return nil
		}
	}

	resp, err := sess.GetNavigationTree(category)
	if err != nil {
		return nil
	}

	for _, n := range resp.Categories.Navnodes {
		fmt.Printf("- %s %d %d\n", n.Name, n.Category.ModelsCount, n.Category.OffersCount)
		for _, nn := range n.Navnodes {
			fmt.Printf("-- %s %d %d nid=%v catid=%v hid=%v target=%s\n",
				nn.Name,
				nn.Category.ModelsCount,
				nn.Category.OffersCount,
				nn.Link.Params.Nid,
				nn.Link.Params.CatID,
				nn.Link.Params.Hid,
				nn.Link.Target,
			)

			totalCount := nn.Category.ModelsCount
			if totalCount == 0 {
				totalCount = nn.Category.OffersCount
			}
			if totalCount > 500000 {
				log.Printf("-- %s skipped (%d offers). For categories with many offers yamarket resp > 10-50sec  ", n.Name, totalCount)
				continue
			}

			for page := 0; page < totalCount/16; page++ {
				w.jobs <- Job{
					Nid:   nn.Link.Params.Nid,
					CatID: nn.Link.Params.CatID,
					Hid:   nn.Link.Params.Hid,
					Page:  page + 1,
				}
				w.postedJobs++
			}
		}
	}
	return sess
}

func (w *YMWorkers) run() {
	var sess *YMSession

	for job := range w.jobs {
		for {
			sess = w.processJob(job, sess)
			if sess != nil {
				break
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}
	w.wg.Done()
	d := atomic.AddInt32(&w.doneWorkers, 1)
	log.Printf("%d workers done\n", d)
}

func (w *YMWorkers) processJob(job Job, sess *YMSession) *YMSession {

	defer func() {
		if v := recover(); v != nil {
			log.Print(v.(error).Error())
			return
		}
	}()

	if sess == nil {
		sess = NewYMSession(proxyList.Next().client, uaList.Next())
		if err := sess.Start(); err != nil {
			return nil
		}
	}

	failedCacheKey := fmt.Sprintf("%v%v%v", job.Nid, job.CatID, job.Hid)
	w.lock.RLock()
	failedPage, found := w.failedQ[failedCacheKey]
	w.lock.RUnlock()
	if found && failedPage < job.Page {
		// This query already finished/failed with early page. skip job
		return sess
	}

	resp, err := sess.GetSearchResults(job.Nid, job.CatID, job.Hid, job.Page)

	if err != nil {
		return nil
	}

	if len(resp.Data.Collections.Products) != 0 {

		for uid, item := range resp.Data.Collections.Products {
			fmt.Printf("P: %s -> min=%s,max=%s\n", item.Titles.Raw, item.Prices.Min, item.Prices.Max)
			rxRepo.PutProdct(item, uid)
		}
	} else if len(resp.Data.Collections.Offers) != 0 {
		for uid, item := range resp.Data.Collections.Offers {
			fmt.Printf("O: %s -> price=%s %s\n", item.Titles.Raw, item.Prices.RawValue, item.Prices.Value)
			rxRepo.PutOffer(item, uid)
		}
	} else {
		log.Printf("%s,%s,%s,%d ->empty\n", job.Nid, job.CatID, job.Hid, job.Page)
		w.lock.Lock()
		w.failedQ[failedCacheKey] = job.Page
		w.lock.Unlock()
	}
	return sess
}
