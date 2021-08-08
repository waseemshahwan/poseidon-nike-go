package Requester

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	cbrotli "github.com/andybalholm/brotli"
	tls "github.com/refraction-networking/utls"
	"golang.org/x/net/http2"
	"poseidon.nike/Fingerprints"
)

type Method int

const (
	GET Method = iota
	POST
	PUT
	PATCH
	DELETE
	NONE Method = -1
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
	Key   string
	Value string
}

type Headers []Header

func(h *Headers) GetHeader(key string) *Header {
	for _, header := range *h {
		if strings.EqualFold(header.Key, key) {
			return &header
		}
	}

	return nil
}

type Request struct {
	Method  Method
	Uri     string
	Body    []byte
	Headers Headers
}

type RawResponse struct {
	Response []byte
	Request  *Request
}

type Response struct {
	Body        []byte
	Status      int
	StatusText  string
	HTTPVersion string
	Headers     Headers
}

type Client struct {
	Request Request
	Channel *chan *RawResponse
}

func ReadGzip(data []byte) ([]byte, error) {
	reader := bytes.NewReader(data)
	gzReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}

	output, err := ioutil.ReadAll(gzReader)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func ReadBr(data []byte) ([]byte, error) {
	byteReader := bytes.NewReader(data)
	reader := cbrotli.NewReader(byteReader)

	output, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func ReadZlib(data []byte) ([]byte, error) {
	reader := bytes.NewReader(data)
	gzReader, err := zlib.NewReader(reader)
	if err != nil {
		return nil, err
	}

	output, err := ioutil.ReadAll(gzReader)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func ValidateRequest(r *Request) error {
	if r.Method > 4 {
		return fmt.Errorf("invalid 'Method' field")
	}

	if len(r.Uri) == 0 {
		return fmt.Errorf("missing 'Uri' field")
	}

	return nil
}

func ConnectTLS(dial net.Conn, config *tls.Config, handshake bool) (*tls.UConn, error) {
	utlsConn := tls.UClient(dial, config, tls.HelloCustom)

	err := utlsConn.ApplyPreset(Fingerprints.FIREFOX_PLAYWRIGHT())
	if err != nil {
		return nil, err
	}

	if handshake == true {
		err := utlsConn.Handshake()
		if err != nil {
			return nil, err
		}
	}

	return utlsConn, nil
}

func HTTP2(request *http.Request, connection net.Conn) (*http.Response, error) {
	request.Proto = "HTTP/2.0"
	request.ProtoMajor = 2
	request.ProtoMinor = 0

	tr := http2.Transport{}
	cConn, err := tr.NewClientConn(connection)
	if err != nil {
		return nil, err
	}

	return cConn.RoundTrip(request)
}

func HTTP1(request *http.Request, connection net.Conn) (*http.Response, error) {
	request.Proto = "HTTP/1.1"
	request.ProtoMajor = 1
	request.ProtoMinor = 1

	err := request.Write(connection)

	if err != nil {
		return nil, err
	}

	return http.ReadResponse(bufio.NewReader(connection), request)
}

func (client *Client) Send() (*RawResponse, error) {

	err := ValidateRequest(&client.Request)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(client.Request.Uri)
	if err != nil {
		return nil, err
	}

	usingTLS := false

	switch u.Scheme {
	case "https":
		usingTLS = true
		if u.Port() == "" {
			u.Host = net.JoinHostPort(u.Host, "443")
		}
	case "http":
		if u.Port() == "" {
			u.Host = net.JoinHostPort(u.Host, "80")
		}
	case "":
		return nil, errors.New("scheme is not specified (http:// or https://)")
	default:
		return nil, errors.New("scheme " + u.Scheme + " is not supported")
	}

	NETConn, err := net.DialTimeout("tcp", u.Host, time.Duration(15)*time.Second)

	var TLSConn *tls.UConn
	if usingTLS {
		config := tls.Config{
			ServerName: u.Hostname(),
			NextProtos: []string{"h2", "http/1.1"},
		}

		TLSConn, err = ConnectTLS(NETConn, &config, true)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(client.Request.Method.String(), u.String(), bytes.NewBuffer(client.Request.Body))
	if err != nil {
		return nil, err
	}

	for _, header := range client.Request.Headers {
		request.Header.Set(header.Key, header.Value)
	}

	request.Close = true
	defer request.Body.Close()

	var response_obj *http.Response
	var response_err error

	if usingTLS {
		//println(TLSConn.ConnectionState().NegotiatedProtocol)
		defer TLSConn.Close()
		switch TLSConn.ConnectionState().NegotiatedProtocol {
		case "h2":
			response_obj, response_err = HTTP2(request, TLSConn)
		case "http/1.1", "":
			response_obj, response_err = HTTP1(request, TLSConn)
		default:
			panic(fmt.Errorf("ALPN not supported"))
		}
	} else {
		if NETConn != nil {
			response_obj, response_err = HTTP1(request, NETConn)
		}
	}

	if response_err != nil {
		return nil, response_err
	}

	resp, err := httputil.DumpResponse(response_obj, true)
	if err != nil {
		println(err)
		return nil, err
	}

	result := &RawResponse{
		Response: resp,
		Request:  &(client.Request),
	}

	if client.Channel != nil {
		*(client.Channel) <- result
	}

	return result, nil

}

func (r *RawResponse) Understand() (*Response, error) {

	response := Response{
		Headers: []Header{},
	}

	var payload string = string(r.Response)
	lines := strings.Split(payload, "\r\n")

	var body []byte
	carriage := 0
	for i, line := range lines {
		if i == 0 {
			response.HTTPVersion = strings.Split(line, " ")[0]
			var err error
			response.Status, err = strconv.Atoi(strings.Split(line, " ")[1])
			if err != nil {
				return nil, err
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
				body = []byte(payload)
				break
			}

			keyIdx := strings.Index(line, ":")
			if keyIdx != -1 {
				key := line[:keyIdx]
				value := strings.TrimSpace(line[keyIdx+1:])
				header := Header{
					Key:   key,
					Value: value,
				}

				response.Headers = append(response.Headers, header)
			}
		}
	}

	if len(body) > 0 {
		// Check if its gzip encoding
		// Check if its chunked or not

		
		var encoding string
		contentEncodingHeader := response.Headers.GetHeader("content-encoding")
		if contentEncodingHeader != nil {
			encoding = strings.ToLower(contentEncodingHeader.Value)
		}
		
		
		chunked := false
		transferEncodingHeader := response.Headers.GetHeader("transfer-encoding")
		if transferEncodingHeader != nil && strings.EqualFold(transferEncodingHeader.Value, "chunked") {
			chunked = true
		}

		decode := func(encoded []byte) ([]byte, error) {
			switch(encoding) {
			case "gzip":
				return ReadGzip(encoded)
			case "deflate":
				return ReadZlib(encoded)
			case "br":
				return ReadBr(encoded)
			default:
				return encoded, nil
			}
		}

		if chunked == true {
			encoded := body
			body = []byte{}
			var size []byte
			for i := 0; i < len(encoded); {
				if encoded[i] == byte(13) && encoded[i+1] == byte(10) {
					sizeDec, err := strconv.ParseInt(string(size), 16, 32)
					if err != nil {
						return nil, err
					}

					if sizeDec == 0 {
						break
					}

					size = []byte{}

					// Get from encoding of i + 2(13, 10) to i + 2(13, 10) + the size of the payload were reading
					data := encoded[i + 2 : i + 2 + int(sizeDec)]

					extracted, err := decode(data)
					if err != nil {
						println("Gzip failed on chunked encoding")
						panic(err)
					}

					body = append(body, extracted...)

					// Set cursor to i + 4(13,10 after size and 13,10 after payload) + the size of the payload we just read
					i =  i + 4 + int(sizeDec)
				} else {
					size = append(size, encoded[i])
					i += 1 // move cursor past to next size byte
				}
			}
		} else {
			var err error
			body, err = decode(body)
			if err != nil {
				return nil, err
			}
		}
		
	}

	response.Body = body

	return &response, nil
}
