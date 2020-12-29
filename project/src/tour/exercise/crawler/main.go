package main

import (
	"fmt"
	"sync"
)

var (
	store = NewStore()
)

type Store struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewStore() *Store {
	return &Store{store: make(map[string]string)}
}

func (s *Store) Set(key, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = val
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.store[key]
	return val, ok
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, wg)
	wg.Wait()
}

func Crawl(url string, depth int, fetcher Fetcher, group *sync.WaitGroup) {
	if depth <= 0 {
		group.Done()
		return
	}
	if _, ok := store.Get(url); ok {
		fmt.Printf("This url (%s) has already fetched... return\n", url)
		group.Done()
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err) // not found.
		group.Done()
		return
	}

	store.Set(url, body)
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		group.Add(1)
		go Crawl(u, depth-1, fetcher, group)
	}
	group.Done()
	return
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
