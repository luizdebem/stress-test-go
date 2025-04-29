package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
	Duration   time.Duration
	Error      error
}

type Report struct {
	TotalTime     time.Duration
	TotalRequests int
	StatusCodes   map[int]int
}

func main() {
	url := flag.String("url", "", "URL a ser testada")
	requests := flag.Int("requests", 0, "Número de requisições a serem realizadas")
	concurrency := flag.Int("concurrency", 1, "Concorrência das requisições")
	flag.Parse()

	if *url == "" || *requests == 0 {
		fmt.Println("URL and number of requests are required")
		return
	}

	results := make(chan Result, *requests)
	var wg sync.WaitGroup
	startTime := time.Now()

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go worker(*url, &wg, results, *requests / *concurrency)
	}

	if remainder := *requests % *concurrency; remainder != 0 {
		wg.Add(1)
		go worker(*url, &wg, results, remainder)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	report := Report{
		StatusCodes: make(map[int]int),
	}

	for result := range results {
		report.TotalRequests++
		report.StatusCodes[result.StatusCode]++
	}

	report.TotalTime = time.Since(startTime)
	printReport(report)
}

func worker(url string, wg *sync.WaitGroup, results chan<- Result, requests int) {
	defer wg.Done()
	client := &http.Client{Timeout: 10 * time.Second}

	for i := 0; i < requests; i++ {
		start := time.Now()
		resp, err := client.Get(url)
		duration := time.Since(start)

		if err != nil {
			results <- Result{Error: err}
			continue
		}
		results <- Result{StatusCode: resp.StatusCode, Duration: duration}
		resp.Body.Close()
	}
}

func printReport(r Report) {
	fmt.Printf("\nSTRESS TEST\n\n")
	fmt.Printf("Tempo de execução: %v\n", r.TotalTime)
	fmt.Printf("Requests realizadas: %d\n", r.TotalRequests)
	fmt.Printf("Total de status 200: %d\n", r.StatusCodes[200])
	fmt.Printf("\nOutros status:\n")
	for code, count := range r.StatusCodes {
		fmt.Printf("  HTTP %d: %d\n", code, count)
	}
}
