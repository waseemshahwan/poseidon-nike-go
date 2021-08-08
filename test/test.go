package main

import (
	"strconv"
	"time"

	"poseidon.nike/Requester"
)

func req(status chan string) {

	headers := []Requester.Header{
		{ Key: "user-agent", Value: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:87.0) Gecko/20100101 Firefox/87.0",},
		{ Key: "accept", Value: "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",},
		{ Key: "accept-language", Value: "en-US,en;q=0.5",},
		{ Key: "accept-encoding", Value: "gzip, deflate, br",},
		{ Key: "upgrade-insecure-requests", Value: "1",},
		{ Key: "connection", Value: "keep-alive",},
		{ Key: "host", Value: "www.nike.com",},
	}

	request := Requester.Request{
		Method: Requester.GET,
		Uri:    "https://www.nike.com/snkrs",
		Headers: headers,
	}

	client := Requester.Client{
		Request: request,
	}

	rawResponse, err := client.Send()
	if err != nil {
		panic(err)
	}

	response, err := rawResponse.Understand()
	if err != nil {
		panic(err)
	}

	status <- strconv.Itoa(response.Status)

}

func main() {
	status := make(chan string)

	go req(status)
	println(<-status)

	time.Sleep(time.Duration(10) * time.Second)
	
	go req(status)
	println(<-status)
}