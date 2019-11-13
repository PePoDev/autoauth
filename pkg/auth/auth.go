package auth

import (
	"sync"
	"time"

	"github.com/pepodev/xlog"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

var globalState *sync.WaitGroup

// StartAutoLogin ...
func StartAutoLogin() {
	globalState = &sync.WaitGroup{}
	globalState.Add(1)
	go func() {
		for true {
			if !IsHeatbeatAlive() {
				Login()
			}
			time.Sleep(time.Second * viper.GetDuration("autoauth.heartbeat.interval"))
		}
	}()
	globalState.Wait()
}

// StopAutoLogin ...
func StopAutoLogin() {
	globalState.Done()
}

// Login will create request to authentication service
func Login() {

}

// Logout ...
func Logout() {

}

// IsHeatbeatAlive ...
func IsHeatbeatAlive() bool {
	code, _, err := fasthttp.GetTimeout(nil, viper.GetString("autoauth.heartbeat.endpoint"), time.Second*viper.GetDuration("autoauth.heartbeat.timeout"))
	xlog.Debugf("Status Code: %v Error Message: %v", code, err)
	if err != nil {
		xlog.Errorf("Heartbeat to %s", viper.GetString("autoauth.heartbeat.endpoint"))
		return false
	}
	xlog.Infof("Heartbeat to %s", viper.GetString("autoauth.heartbeat.endpoint"))
	return true
}
