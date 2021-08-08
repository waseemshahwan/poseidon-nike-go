package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

const URL = "https://mpsnare.iesnare.com/snare.js"

func main() {
	u, err := url.Parse(URL)
	if err != nil {
		panic(err)
	}

	host := u.Host

	splitHost := strings.Split(host, ":")
	hostname := splitHost[0]
	port := ""

	if len(splitHost) == 2 {
		port = splitHost[1]
	}

	if len(port) == 0 {
		if strings.Contains(URL, "https") {
			port = "443"
		} else {
			port = "80"
		}

		host = hostname + ":" + port
	}

	config := tls.Config{ServerName: hostname}
	dialConn, err := net.DialTimeout("tcp", host, time.Duration(15)*time.Second)

	time.Sleep(time.Duration(2) * time.Second)

	tlsConn := tls.Client(dialConn, &config)

	time.Sleep(time.Duration(2) * time.Second)

	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Duration(2) * time.Second)

	request.Proto = "HTTP/2.0"
	request.ProtoMajor = 2
	request.ProtoMinor = 0

	res, err := httputil.DumpRequest(request, true)

	time.Sleep(time.Duration(2) * time.Second)

	a, err := tlsConn.Write(res)
	println(a, err)

	time.Sleep(time.Duration(2) * time.Second)

	time.Sleep(time.Duration(5) * time.Second)

	var b []byte
	tlsConn.Read(b)

	println(b)

	//tlsConn.Write(request.)
	/*
		client.RoundTrip(request)

		r, err := httputil.DumpResponse(res, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(r)
	*/
}
