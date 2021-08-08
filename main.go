package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	pw "github.com/mxschmitt/playwright-go"
	"poseidon.nike/Requester"
)

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

var URL string = "https://tlsfingerprint.io"

func main() {

	pwe, err := pw.Run()
	checkPanic(err)

	options := pw.BrowserTypeLaunchOptions{}
	options.Headless = &[]bool{false}[0]

	browser, err := pwe.Firefox.Launch(options)
	checkPanic(err)

	page, err := browser.NewPage()
	checkPanic(err)

	//time.Sleep(time.Second * time.Duration(100000))

	logName := "logs/" + URL[strings.Index(URL, "//")+2:] + "-" + time.Now().Format("15-04-05") + ".log"
	fmt.Println("Printing logs to ", logName)
	f, err := os.Create(logName)
	if err != nil {
		println("opening file failed")
	}

	defer f.Close()

	page.Route("*", func(route pw.Route, req pw.Request) {

		headers := []Requester.Header{}
		realHeaders := req.Headers()
		for key, value := range realHeaders {
			headers = append(headers, Requester.Header{
				Key:   key,
				Value: value,
			})
		}

		body, err := req.PostDataBuffer()
		checkPanic(err)

		request := Requester.Request{
			Uri:     req.URL(),
			Method:  Requester.MethodFrom(req.Method()),
			Headers: headers,
		}

		whitelistPrefixes := []string{
			"http://",
			"https://",
			"http://www.",
			"https://www.",
		}

		whitelistUrls := []string{
			"example.com",
			"api.nike.com",
			"unite.nike.com",
		}

		for _, i := range whitelistUrls {
			for _, k := range whitelistPrefixes {
				if strings.Index(req.URL(), k+i) == 0 || strings.Index(req.URL(), k+i) == 0 {
					url := req.URL()
					if len(url) > 100 {
						url = url[:100]
					}

					fmt.Println(Requester.MethodFrom(req.Method()).String(), url)
					break
				}
			}
		}

		if len(body) > 0 {
			request.Body = body
		}

		client := Requester.Client{
			Request: request,
		}

		rawResponse, err := client.Send()
		if err != nil {
			fmt.Println(req.Method(), req.URL())
			fmt.Println("ABORT: ", err)
			route.Abort()
			return
		}

		response, err := rawResponse.Understand()
		if err != nil {
			fmt.Println(req.Method(), req.URL())
			fmt.Println("ABORT: ", err)
			route.Abort()
			return
		}

		var headersStr string
		headersRes := map[string]string{}
		for k, headerRes := range response.Headers {
			headersRes[headerRes.Key] = headerRes.Value
			headersStr += "\t" + headerRes.Key + ": " + headerRes.Value
			if k != len(response.Headers)-1 {
				headersStr += "\n"
			}
		}

		f.Write([]byte("\n\n\n---- Request ----\nURL: " + req.URL() + "\nMethod: " + req.Method() + "\nHTTPVersion: " + response.HTTPVersion + "\n\n--- Response ---\nStatus: " + strconv.Itoa(response.Status) + "\nHeaders:\n" + headersStr + "\nBody:\n" + string(response.Body) + "\n--- END ---\n\n\n"))

		ro := pw.RouteFulfillOptions{
			Status:  &response.Status,
			Body:    response.Body,
			Headers: headersRes,
		}

		route.Fulfill(ro)

	})

	//_, err = page.Goto("https://nike.com/snkrs")
	// _, err = page.Goto("https://client.tlsfingerprint.io:8443/")
	_, err = page.Goto(URL)
	checkPanic(err)

	time.Sleep(time.Second * time.Duration(100000))

}
