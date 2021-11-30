package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	v  map[string]bool
}

func (c *SafeMap) SetVal(key string, val bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.v[key] = val
}

func (c *SafeMap) GetVal(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	_, ok := c.v[key]
	return ok
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, urlMap *SafeMap, ch chan bool) {
	if depth <= 0 {
		ch <- true
		return
	}
	if !urlMap.GetVal(url) {
		urlMap.SetVal(url, true)
	} else {
		ch <- true
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("err", err)
		urlMap.SetVal(url, false)
		ch <- false
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	routines := make([]chan bool, len(urls))
	for index, url := range urls {
		routines[index] = make(chan bool)
		go Crawl(url, depth-1, fetcher, urlMap, routines[index])
	}

	for _, childRoutine := range routines {
		<-childRoutine
	}

	ch <- true
	return
}

func main() {
	ch := make(chan bool)
	depth := 4
	urlMap := SafeMap{v: make(map[string]bool)}
	go Crawl("https://golang.org/", depth, fetcher, &urlMap, ch)
	<-ch
	fmt.Println("main finished")
}


type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}


var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
