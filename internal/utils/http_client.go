package utils

import (
	"strings"

	"github.com/pepodev/xlog"
	"github.com/valyala/fasthttp"
)

// Do will send http request and return the response
func Do(url string, method string, headers []string, bodys []string) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	for _, header := range headers {
		h := strings.Split(header, ":")
		req.Header.Add(h[0], h[1])
	}

	if strings.ToUpper(method) != "GET" {
		for _, body := range bodys {
			req.AppendBody([]byte(body))
		}
	}

	req.Header.SetMethod(strings.ToUpper(method))
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)

	xlog.Debugf("[code: %v] [body: %v] [err: %v]", resp.StatusCode(), string(resp.Body()), err)
	if err != nil {
		return err
	}
	return nil
}
