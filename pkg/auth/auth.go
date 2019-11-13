package auth

import (
	"sync"
	"time"

	"github.com/pepodev/autoauth/internal/utils"

	"github.com/pepodev/xlog"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

var globalState *sync.WaitGroup

var (
	heartbeatEndpoint string
	heartbeatTimeout  time.Duration
)

// StartAutoLogin ...
func StartAutoLogin() {
	heartbeatEndpoint = utils.ViperGetString("autoauth.heartbeat.endpoint")
	heartbeatTimeout = utils.ViperGetDuration("autoauth.heartbeat.timeout")

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
	code, _, err := fasthttp.GetTimeout(nil, heartbeatEndpoint, time.Second*heartbeatTimeout)
	xlog.Debugf("code: %v err: %v", code, err)
	if err != nil {
		xlog.Errorf("Heartbeat to %s", heartbeatEndpoint)
		return false
	}
	xlog.Infof("Heartbeat to %s", heartbeatEndpoint)
	return true
}
