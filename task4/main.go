package main

import (
	"fmt"
	"sync"
)

var safeMap = SomeSafeStruct{visited: make(map[string]bool)}
var wg *sync.WaitGroup

type SomeSafeStruct struct {
	mux     sync.Mutex
	visited map[string]bool
}

func (s *SomeSafeStruct) Add(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.visited[key] = true
}

func (s *SomeSafeStruct) Contains(key string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.visited[key]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	// Чекаем если есть
	if safeMap.Contains(url) {
		return
	}

	// check if there is an error fetching
	body, urls, err := fetcher.Fetch(url)
	safeMap.Add(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// list contents and crawl recursively
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1) //  Чтобы это говно не закрывалос раньше времени

		go Crawl(u, depth-1, fetcher)
	}
}

func main() {
	wg = &sync.WaitGroup{}
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	//time.Sleep(time.Second)   не надо так делать наверное
	wg.Wait()
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
