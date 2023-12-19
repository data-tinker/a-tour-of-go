package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type UrlChecker interface {
	Set(url string)
	Check(url string) bool
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, checker UrlChecker, wg *sync.WaitGroup) {
	// Fetch URLs in parallel.
	// Don't fetch the same URL twice.
	defer wg.Done()

	if depth <= 0 {
		return
	}

	if checker.Check(url) {
		return
	}

	checker.Set(url)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	var wg_ sync.WaitGroup
	for _, u := range urls {
		wg_.Add(1)
		go Crawl(u, depth-1, fetcher, checker, &wg_)
	}

	wg_.Wait()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	checker := urlChecker{visited: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, &checker, &wg)
	wg.Wait()
}

type urlChecker struct {
	visited map[string]bool
	mut     sync.Mutex
}

func (checker *urlChecker) Set(url string) {
	checker.mut.Lock()
	defer checker.mut.Unlock()
	checker.visited[url] = true
}

func (checker *urlChecker) Check(url string) bool {
	checker.mut.Lock()
	defer checker.mut.Unlock()
	_, ok := checker.visited[url]
	return ok
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
