package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Func func(key string, done ...chan<- struct{}) (interface{}, error)

type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	requests chan request
}

type entry struct {
	res   result
	ready chan struct{}
}

func New(f Func) *Memo {
	memo := &Memo{
		requests: make(chan request),
	}
	go memo.server(f)
	return memo
}

func (m *Memo) Get(key string, done ...<-chan struct{}) (interface{}, error) {
	response := make(chan result)
	m.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (m *Memo) Close() { close(m.requests) }

func (m *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range m.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(resp chan<- result) {
	<-e.ready
	resp <- e.res
}

func httpGetBody(url string, done ...chan<- struct{}) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func incomingURLS() []string {
	return []string{
		"https://www.stackoverflow.com",
		"https://go.dev",
		"https://www.paypal.com",
		"https://www.pornhub.com",
		"https://www.whattomine.com",
		"https://www.4chan.org",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.tiktok.com",
		"https://www.instagram.com",
		"https://www.gamespot.com",
		"https://www.glovoapp.com",
		"https://www.donesi.com",
	}
}

func main() {
	m := New(httpGetBody)
	defer m.Close()
	var wg sync.WaitGroup
	for _, url := range incomingURLS() {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			val, err := m.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(val.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}
