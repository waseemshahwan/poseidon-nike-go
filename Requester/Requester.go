package Requester

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	tls "github.com/refraction-networking/utls"
	"golang.org/x/net/http2"
	"poseidon.nike/DataStream"
)

type Method int

const (
	GET Method = iota
	POST
	PUT
	PATCH
	DELETE
)

func MethodFrom(d string) Method {
	methods := [...]string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	for k, method := range methods {
		if strings.EqualFold(d, method) {
			return Method(k)
		}
	}

	panic(fmt.Errorf("method does not exist"))
}

func (m Method) String() string {
	methods := [...]string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	if len(methods) < int(m) {
		panic(fmt.Errorf("invalid 'Method' type"))
	}
 
	return methods[m]
}

type Header struct {
	Key		string
	Value	string
}

type Request struct {
	Method		Method
	Uri			string
	Body		*DataStream.DataStream
	Headers		[]Header
}

type RawResponse struct {
	Response	*DataStream.DataStream
	Request		*Request
}

type Response struct {
	Body		[]byte
	Status		int
	StatusText	string
	HTTPVersion	string
	Headers		[]Header
}

type Client struct {
	Request		Request
}

func ValidateRequest(r *Request) {
	if r.Method > 4 { panic(fmt.Errorf("invalid 'Method' field")) }
	if len(r.Uri) == 0 { panic(fmt.Errorf("missing 'Uri' field")) }
}

func(r *RawResponse) Understand() Response {

	response := Response{
		Headers: []Header{},
	}

	buf := new([]byte)
	buf_sz := 0
	r.Response.Read(buf, &buf_sz)

	var payload string = string(*buf)
	lines := strings.Split(payload, "\r\n")

	carriage := 0
	for i, line := range lines {
		if i == 0 {
			response.HTTPVersion = strings.Split(line, " ")[0]
			var err error
			response.Status, err = strconv.Atoi(strings.Split(line, " ")[1])
			if err != nil {
				panic(err)
			}

			response.StatusText = strings.Join(strings.Split(line, " ")[2:], " ")
		} else {

			if line == "" {
				carriage++
			}

			if carriage > 0 {
				payload := strings.Join(lines[i:], "\r\n")
				if payload[0] == 13 && payload[1] == 10 { // \r\n
					payload = payload[2:]
				}
				response.Body = []byte(payload)
				break
			}

			keyIdx := strings.Index(line, ":")
			if keyIdx != -1 {
				key := line[:keyIdx]
				value := strings.TrimSpace(line[keyIdx+1:])
				header := Header{
					Key: key,
					Value: value,
				}

				response.Headers = append(response.Headers, header)
			}
		}
	}

	return response
}

func(client *Client) Send() RawResponse {
	ValidateRequest(&client.Request)

	u, err := url.Parse(client.Request.Uri)
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
		if strings.Contains(client.Request.Uri, "https") {
			port = "443"
		} else {
			port = "80"
		}

		host = hostname + ":" + port
	}

	config := tls.Config { ServerName: hostname }
	dialConn, err := net.DialTimeout("tcp", host, time.Duration(15) * time.Second)

	if err != nil {
		panic(err)
	}

	utlsConn := tls.UClient(dialConn, &config, tls.HelloFirefox_63)
	defer utlsConn.Close()

	err = utlsConn.Handshake()
	if err != nil {
		 panic(fmt.Errorf("uTlsConn.Handshake() error: %+v", err))
	}

	buf := new([]byte)
	buf_sz := 0
	if client.Request.Method.String() != "GET" {
		client.Request.Body.Read(buf, &buf_sz)
	}

	request, err := http.NewRequest(client.Request.Method.String(), u.String(), bytes.NewBuffer(*buf))
	if err != nil {
		panic(err)
	}

	var response_obj *http.Response
	var response_err error

	switch utlsConn.HandshakeState.ServerHello.AlpnProtocol {
	case "h2":
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

	ds := DataStream.New()
	ds.Write(&resp)
	ds.Close()

	return RawResponse{
		Response: ds,
		Request: &client.Request,
	}
	
}