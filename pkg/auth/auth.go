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
	heartbeatInterval time.Duration
)

// StartAutoLogin ...
func StartAutoLogin() {
	heartbeatEndpoint = utils.ViperGetString("autoauth.heartbeat.endpoint")
	heartbeatTimeout = utils.ViperGetDuration("autoauth.heartbeat.timeout")
	heartbeatInterval = viper.GetDuration("autoauth.heartbeat.interval")

	globalState = &sync.WaitGroup{}
	globalState.Add(1)
	go func() {
		for true {
			if !IsHeatbeatAlive() {
				Login()
			}
			time.Sleep(time.Second * heartbeatInterval)
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
		xlog.Errorf("Heartbeat to %s is OK", heartbeatEndpoint)
		return false
	}
	xlog.Infof("Heartbeat to %s", heartbeatEndpoint)
	return true
}
