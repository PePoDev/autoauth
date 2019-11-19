package utils

import (
	"crypto/tls"
	"net"
	"strings"
	"time"

	"github.com/pepodev/xlog"
	"github.com/valyala/fasthttp"
)

var client *fasthttp.Client

func init() {
	client = &fasthttp.Client{
		ReadTimeout:                   10 * time.Second,
		DisableHeaderNamesNormalizing: false,
		MaxConnsPerHost:               233,
		MaxIdleConnDuration:           100 * time.Second,
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Second*time.Duration(10))
		},
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
}

// Do will send http request and return the response
func Do(url string, method string, headers []string, bodys []string, timeout time.Duration) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	var tempHeader []string
	for _, header := range headers {
		tempHeader = strings.Split(header, ":")
		req.Header.Add(tempHeader[0], tempHeader[1])
	}

	var requestBodyString string
	if strings.ToUpper(method) != "GET" {
		for _, body := range bodys {
			requestBodyString += body + "&"
		}
	}
	req.SetBodyString(requestBodyString)

	xlog.Debugf("req body: %s", string(req.Body()))

	req.Header.SetMethod(strings.ToUpper(method))
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()

	err := client.DoTimeout(req, resp, timeout)
	xlog.Debugf("[code: %v] [body: %v] [err: %v]", resp.StatusCode(), string(resp.Body()), err)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
