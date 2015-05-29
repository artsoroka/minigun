package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	log.Println("Program started...")

	rounds := flag.Int("rounds", 5, "number of requests")
	url := flag.String("url", "http://example.com", "url address")
	file := flag.String("file", "null", "path to file with request body")

	flag.Parse()

	if *file == "null" {
		panic("no file is provided")
	}

	reqBody, err := ioutil.ReadFile(*file)

	if err != nil {
		panic("Something wrong with file")
	}

	channel := make(chan string, *rounds)
	for i := 0; i < *rounds; i++ {
		go func(url string) {
			log.Println("connecting " + url)
			log.Println(url)

			resp, err := http.Post(url, "text/xml", bytes.NewBuffer(reqBody))
			if err != nil {
				panic(err)
			}
			channel <- resp.Status
		}(*url)
	}
	for i := 0; i < *rounds; i++ {

		status := <-channel
		log.Printf("Response status: " + status)
	}
	elapsed := time.Since(start)
	log.Printf("Time taken %s", elapsed)
}
