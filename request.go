package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/refraction-networking/utls"
	"golang.org/x/net/http2"
)

func MakeRequest(uri string, method string) {

	u, err := url.Parse(uri)
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

	fmt.Println(host, hostname, port)

	if len(port) == 0 {
		if strings.Contains(uri, "https") {
			port = "443"
		} else {
			port = "80"
		}

		host = hostname + ":" + port
	}

	fmt.Println(host, hostname, port)

	config := tls.Config { ServerName: hostname }
	dialConn, err := net.DialTimeout("tcp", host, time.Duration(15) * time.Second)

	if err != nil {
		panic(err)
	}

	utlsConn := tls.UClient(dialConn, &config, tls.HelloChrome_72)
	defer utlsConn.Close()

	err = utlsConn.Handshake()
	if err != nil {
		 panic(fmt.Errorf("uTlsConn.Handshake() error: %+v", err))
	}

	request := &http.Request{
		Method: method,
		URL: u,
		Header: make(http.Header),
		Host: hostname,
	}

	var response_obj *http.Response
	var response_err error

	switch utlsConn.HandshakeState.ServerHello.AlpnProtocol {
	case "h2":
		fmt.Println("test-1")
		request.Proto = "HTTP/2.0"
		request.ProtoMajor = 2
		request.ProtoMinor = 0

		tr := http2.Transport{}
		cConn, err := tr.NewClientConn(utlsConn)
		if err != nil {
			panic(err)
		}

		response_obj, response_err = cConn.RoundTrip(request)
	case "http/1.1", "":
		fmt.Println("test-2")
		request.Proto = "HTTP/1.1"
		request.ProtoMajor = 1
		request.ProtoMinor = 1

		err := request.Write(utlsConn)

		if err != nil {
			panic(err)
		}

		response_obj, response_err = http.ReadResponse(bufio.NewReader(utlsConn), request)
	default:
		panic(fmt.Errorf("ALPN PROTOCOL NOT SUPPORTED"))
	}

	if response_err != nil {
		panic(response_err)
	}

	resp, err := httputil.DumpResponse(response_obj, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resp))
	
}
