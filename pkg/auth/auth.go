package auth

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pepodev/autoauth/internal/message"
	"github.com/pepodev/autoauth/internal/utils"

	"github.com/pepodev/xlog"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

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

	fmt.Println(message.GetWelcome())
	xlog.Info("AutoAuth Started")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for true {
			if !IsHeatbeatAlive() {
				Login()
			}
			time.Sleep(time.Second * heartbeatInterval)
		}
	}()

	xlog.Infof("os %v AutoAuth stopped", <-sigs)
}

// StopAutoLogin ...
func StopAutoLogin() {

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
