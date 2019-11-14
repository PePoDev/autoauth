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
func Do(url string, method string, headers []string, bodys []string) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	for _, header := range headers {
		h := strings.Split(header, ":")
		req.Header.Add(h[0], h[1])
	}

	temp := ""
	if strings.ToUpper(method) != "GET" {
		for _, body := range bodys {
			temp += body + "&"
		}
	}

	req.SetBodyString(temp)

	xlog.Debugf("resp: %s", string(req.Body()))

	req.Header.SetMethod(strings.ToUpper(method))
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()

	err := client.DoTimeout(req, resp, time.Second*10)
	xlog.Debugf("[code: %v] [body: %v] [err: %v]", resp.StatusCode(), string(resp.Body()), err)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
