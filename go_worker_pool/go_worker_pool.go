package main

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	Status string
	URL    string
}

func crawl(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("Worker no. %d\n", wId)
		res, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		results <- Result{Status: res.Status, URL: site.URL}
	}
}

/*
Worker Pools:
The main reason for using worker pools is to avoid the overhead of spinning up a new thread
every time we have a job.
But the cost of spinning up a goroutine is relatively insignificant compared to the cost of
spinning up a goroutine is relatively insignificant compared to the cost of spinning up a
thread. So the same reasoning dosen't really apply here.

Another reason why worker pool is useful is to limit the number of concurrent processes running
at any point of time.

*/

func main() {
	fmt.Println("Go Worker Pool Example:")

	//create the channels to send jobs to the workers and to receive results from the workers
	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	//initiate the workers in the worker pool
	noOfWorkers := 3
	for wId := 1; wId <= noOfWorkers; wId++ {
		go crawl(wId, jobs, results)
	}

	//these are some site urls to be used as jobs for our workers
	//normally a producer would be continuously producing such jobs for our workers
	urls := []string{
		"https://www.google.com",
		"https://www.example.com",
		"https://www.youtube.com",
		"https://www.cloudflare.com",
		"https://www.amazon.com",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for i := 1; i <= len(urls); i++ {
		result := <-results
		log.Println(result)
	}

}
