package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	const bufferSize = 2

	sites := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.github.com",
		"https://www.linkedin.com",
	}
	urlChannel := make(chan string, bufferSize)
	responseChannel := make(chan string, bufferSize)
	var wg sync.WaitGroup
	wg.Add(4)

	go publishURLs(urlChannel, sites)
	go readURLs(urlChannel, responseChannel)
	go readResponses(responseChannel, &wg)

	wg.Wait()
	fmt.Println("elapsed_time", time.Since(startTime).String())
}

func readResponses(ch chan string, wg *sync.WaitGroup) error {
	for response := range ch {
		fmt.Printf("Received response %s\n", response)
		wg.Done()

	}
	return nil
}

func readURLs(urlChannel chan string, responseChannel chan string) {
	for url := range urlChannel {
		fmt.Println("Url recebida", url)
		go requestURL(url, responseChannel)
	}
}

func publishURLs(urlChannel chan string, sites []string) {
	for _, site := range sites {
		fmt.Println("URL enviada", site)
		urlChannel <- site
	}
	close(urlChannel)
}

func requestURL(url string, responseChannel chan<- string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		responseChannel <- fmt.Sprintf("error creating HTTP request: %s", err)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		responseChannel <- fmt.Sprintf("error making HTTP request: %s", err)
		return
	}
	responseChannel <- res.Status
}
