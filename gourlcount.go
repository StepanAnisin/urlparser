package urlparser

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
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
func handleUrl(request_url string, r chan int) {
	response := getResponceBody(request_url)
	count := getGoWordCount(response)
	r <- count
}

func GetGoWordCountByUrls(urls []string) {
	result := 0
	//Limit goroutines number
	const maxjobs = 5
	r := make(chan int, maxjobs)
	for _, request_url := range urls {
		//validate url
		_, err := url.ParseRequestURI(request_url)
		if err != nil {
			panic(err)
		}
		// blocks
		go handleUrl(request_url, r)
		result += <-r
	}
	close(r)
	fmt.Println(result)
}
