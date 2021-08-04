package main

import (
	"fmt"
	"time"

	pw "github.com/mxschmitt/playwright-go"
	"poseidon.nike/DataStream"
	"poseidon.nike/Requester"
)

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	
	pwe, err := pw.Run()
	checkPanic(err)

	options := pw.BrowserTypeLaunchOptions{}
	options.Headless = &[]bool{false}[0]

	browser, err := pwe.Firefox.Launch(options)
	checkPanic(err)

	page, err := browser.NewPage()
	checkPanic(err)

	page.Route("*", func(route pw.Route, req pw.Request) {
		
		headers := []Requester.Header{}
		realHeaders := req.Headers()
		for  key, value := range realHeaders {
			headers = append(headers, Requester.Header{
				Key: key,
				Value: value,
			})
		}

		body, err := req.PostDataBuffer()
		checkPanic(err)

		request := Requester.Request{
			Uri: req.URL(),
			Method: Requester.MethodFrom(req.Method()),
			Headers: headers,
		}

		fmt.Println(req.URL(), Requester.MethodFrom(req.Method()).String())
		for _, v := range headers {
			fmt.Println(v.Key, v.Value)
		}
		
		if len(body) > 0 {
			bodyStream := DataStream.New()
			bodyStream.Write(&body)
			bodyStream.Close()

			request.Body = bodyStream
		}

		client := Requester.Client{
			Request: request,
		}

		rawResponse := client.Send()
		response := rawResponse.Understand()

		headersRes := map[string]string{}
		for _, headerRes := range response.Headers {
			headersRes[headerRes.Key] = headerRes.Value
		}

		for k, v := range headersRes {
			fmt.Println(k, v)
		}

		ro := pw.RouteFulfillOptions{
			Status: &response.Status,
			Body: response.Body,
			Headers: headersRes,
		}

		route.Fulfill(ro)

	})

	_, err = page.Goto("https://nike.com/snkrs")
	checkPanic(err)

	time.Sleep(time.Second * time.Duration(100000))


}