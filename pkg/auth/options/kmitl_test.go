package options_test

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

func TestKMITL(t *testing.T) {
	code, _, err := fasthttp.GetTimeout(nil, "http://google.com/", 5*time.Second)
	log.Printf("[status: %d] [err: %v]", code, err)
	if err != nil {
		t.Fail()
	}
	if code != http.StatusOK {
		t.Fail()
	}
}
