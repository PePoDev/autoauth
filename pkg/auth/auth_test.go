package auth_test

import (
	"bytes"
	"testing"

	"github.com/pepodev/autoauth/internal/utils"
	"github.com/pepodev/autoauth/pkg/auth"
	"github.com/spf13/viper"
)

func TestHeartbeat(t *testing.T) {
	t.Skip("ignroe this function: waiting improve load preset workflow")
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)
	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	preset := auth.LoadPresetFromPath(utils.GetWorkingDirectory(), "autoauth.yml")
	err := preset.IsHeartbeatAlive()
	if err != nil {
		t.Errorf("Heartbeat err: %v", err)
	}
}
