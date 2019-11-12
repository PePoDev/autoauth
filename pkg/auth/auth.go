package auth

import (
	"net/http"

	"github.com/pepodev/xlog"
	"github.com/valyala/fasthttp"
)

// StartAutoLogin ...
func StartAutoLogin() {

}

// StopAutoLogin ...
func StopAutoLogin() {

}

// Login will create request to authentication service
func Login(option LoginOption) {
	if err := option.Login(); err != nil {
		xlog.Errorf("error: %v", err)
	}
}

// Logout ...
func Logout() {

}

// Heartbeats ...
func Heartbeats() bool {
	code, _, err := fasthttp.GetTimeout(nil, "http://google.com/", 5)
	if err != nil {
		return false
	}
	if code != http.StatusOK {
		return false
	}
	return true
}
