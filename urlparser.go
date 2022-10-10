package urlparser

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sync"
)

// Get url response body
func getResponseBody(request_url string) string {
	resp, err := http.Get(request_url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

// Get "Go" count in url
// It finds exact match
// for example, "Golang, golang" will not count
// but ">Go<" will
func getGoWordCount(value string) int {
	// Match non-space character sequences.
	var containsCanRegex = regexp.MustCompile(`\bGo\b`)
	containsCanRegex.MatchString(value)
	// Find all matches and return count.
	results := containsCanRegex.FindAllString(value, -1)
	return len(results)
}

func handleUrl(request_url string, wg *sync.WaitGroup, block chan struct{}, result *int, mutex *sync.Mutex) {
	defer wg.Done()
	response := getResponseBody(request_url)
	count := getGoWordCount(response)
	fmt.Printf("Count for %s:%d \n", request_url, count)
	mutex.Lock()
	*result += count
	mutex.Unlock()
	// unlock queue
	<-block
}

func GetGoWordCountByUrls(urls []string) int {
	result := 0
	const maxjobs = 5
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(len(urls))
	// channel to limit goroutine number
	block := make(chan struct{}, maxjobs)
	defer close(block)
	for _, request_url := range urls {
		//validate url
		_, err := url.ParseRequestURI(request_url)
		if err != nil {
			panic(err)
		}
		block <- struct{}{}
		go handleUrl(request_url, &wg, block, &result, &mutex)
	}
	wg.Wait()
	fmt.Printf("Total: %d", result)
	return result
}
