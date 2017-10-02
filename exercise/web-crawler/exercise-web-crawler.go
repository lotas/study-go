package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type FetchCache struct {
	urls map[string]bool
	mux  sync.Mutex
}

func (cache *FetchCache) hasURL(url string) bool {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	return cache.urls[url]
}

func (cache *FetchCache) setURL(url string, k bool) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	cache.urls[url] = k
}

var cache = FetchCache{urls: make(map[string]bool)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 || cache.hasURL(url) {
		return
	}

	cache.setURL(url, true)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	// ch <- url

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}

	return
}

var ch chan string

func main() {
	ch = make(chan string)

	// func() {
	Crawl("http://golang.org/", 4, fetcher)
	// close(ch)
	// }()

	// for url := range ch {
	// 	fmt.Println("Fetched url", url)
	// }
	time.Sleep(time.Second)
}

// fakeFetcher is Fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
