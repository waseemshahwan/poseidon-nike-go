package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	buf := make([]byte, 1024)
	for true {
		read, err := req.Body.Read(buf)
		if err != nil {
			panic(err)
		}

		if read > 0 {
			fmt.Print(string(buf[:read]))
		}
	}
}

func main() {

    http.HandleFunc("/", hello)

    http.ListenAndServe(":8090", nil)
}