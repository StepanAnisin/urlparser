package urlparser

import (
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sync"
)

// Get url response body
func getResponceBody(request_url string) string {
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

// This function handles url asynchronously
func handleUrl(request_url string, result *int, mutex *sync.Mutex) {
	response := getResponceBody(request_url)
	count := getGoWordCount(response)
	mutex.Lock()
	*result += count
	mutex.Unlock()
}

func GetGoWordCountByUrls(urls []string) {
	result := 0
	const maxjobs = 5
	var mutex sync.Mutex
	//iterate over urls w/o index
	for _, requset_url := range urls {
		//validate url
		_, err := url.ParseRequestURI(requset_url)
		if err != nil {
			panic(err)
		}
		go handleUrl(requset_url, &result, &mutex)
	}
}
