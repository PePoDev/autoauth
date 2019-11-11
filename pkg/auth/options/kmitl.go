package options

import (
	"fmt"
	"net/http"

	"github.com/pepodev/autoauth/pkg/auth"

	"github.com/valyala/fasthttp"
)

// KMITL ...
type KMITL struct {
	auth.LoginBase
}

// Login ...
func (req KMITL) Login() error {
	code, _, err := fasthttp.GetTimeout(nil, "http://google.com/", 5)
	if err != nil {
		return err
	}
	if code != http.StatusOK {
		return fmt.Errorf("error with status code %v", code)
	}
	return nil
}
