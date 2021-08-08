package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	gotls "crypto/tls"

	utls "github.com/refraction-networking/utls"
	"golang.org/x/net/http2"
)

func main() {
	utlsClientHelloIDName := flag.String("utls", "", "use utls with the given ClientHelloID (e.g. HelloGolang)")
	callHandshake := flag.Bool("callhandshake", false, "call UConn.Handshake inside DialTLS")
	flag.Parse()

	if *callHandshake && *utlsClientHelloIDName == "" {
		fmt.Fprintf(os.Stderr, "error: -callhandshake only makes sense with -utls\n")
		os.Exit(1)
	}

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "error: need a URL\n")
		os.Exit(1)
	}
	url := flag.Arg(0)

	utlsClientHelloID, ok := map[string]*utls.ClientHelloID{
		"":                      nil,
		"HelloGolang":           &utls.HelloGolang,
		"HelloRandomized":       &utls.HelloRandomized,
		"HelloRandomizedALPN":   &utls.HelloRandomizedALPN,
		"HelloRandomizedNoALPN": &utls.HelloRandomizedNoALPN,
		"HelloFirefox_Auto":     &utls.HelloFirefox_Auto,
		"HelloFirefox_55":       &utls.HelloFirefox_55,
		"HelloFirefox_56":       &utls.HelloFirefox_56,
		"HelloFirefox_63":       &utls.HelloFirefox_63,
		"HelloChrome_Auto":      &utls.HelloChrome_Auto,
		"HelloChrome_58":        &utls.HelloChrome_58,
		"HelloChrome_62":        &utls.HelloChrome_62,
		"HelloChrome_70":        &utls.HelloChrome_70,
		"HelloIOS_Auto":         &utls.HelloIOS_Auto,
		"HelloIOS_11_1":         &utls.HelloIOS_11_1,
	}[*utlsClientHelloIDName]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown client hello ID %q\n", *utlsClientHelloIDName)
		os.Exit(1)
	}

	tr := &http2.Transport{}
	if utlsClientHelloID != nil {
		tr.DialTLS = func(network, addr string, cfg *gotls.Config) (net.Conn, error) {
			fmt.Printf("DialTLS(%q, %q)\n", network, addr)

			conn, err := net.Dial(network, addr)
			if err != nil {
				return nil, err
			}

			fmt.Println("Dial[ed] to network: ", network, " & address: ", addr)
			time.Sleep(time.Duration(2) * time.Second)

			uconn := utls.UClient(conn, &utls.Config{NextProtos: cfg.NextProtos}, *utlsClientHelloID)

			println("Initiated TLS Client")
			time.Sleep(time.Duration(2) * time.Second)

			colonPos := strings.LastIndex(addr, ":")
			if colonPos == -1 {
				colonPos = len(addr)
			}

			uconn.SetSNI(addr[:colonPos])

			if *callHandshake {
				err = uconn.Handshake()
				println("Handshake complete")
				time.Sleep(time.Duration(2) * time.Second)
			}

			return uconn, err
		}
	}

	for i := 0; i < 1; i++ {
		resp, err := get(tr, url)
		if err != nil {
			fmt.Printf("%2d err %v\n", i, err)
		} else {
			fmt.Printf("%2d %s %s\n", i, resp.Proto, resp.Status)

			data, err := httputil.DumpResponse(resp, true)
			if err != nil {
				panic(err)
			}

			fmt.Println(len(data))
		}
	}
}

func get(rt http.RoundTripper, url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Request created")
	time.Sleep(time.Duration(2) * time.Second)

	resp, err := rt.RoundTrip(req)
	if err != nil {
		println("broken")
		return nil, err
	}

	println("Round trip complete")
	time.Sleep(time.Duration(2) * time.Second)

	// Read and close the body to enable connection reuse with HTTP/1.1.
	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
