package main

import (
	"fmt"
	"sync"
)

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	var execute func(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup)
	execute = func(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
		if depth <= 0 {
			wg.Done()
			return
		}
		if _, ok := store.Get(url); ok {
			fmt.Printf("This url (%s) has already fetched... return\n", url)
			wg.Done()
			return
		}
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err) // not found.
			wg.Done()
			return
		}

		store.Set(url, body)
		fmt.Printf("found: %s %q\n", url, body)

		for _, u := range urls {
			wg.Add(1)
			go execute(u, depth-1, fetcher, wg)
		}
		wg.Done()
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go execute(url, depth, fetcher, wg)
	wg.Wait()
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

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
