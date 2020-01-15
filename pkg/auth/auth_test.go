package auth_test

import (
	"testing"

	"github.com/pepodev/autoauth/internal/utils"
	"github.com/pepodev/autoauth/pkg/auth"
)

func TestHeartbeat(t *testing.T) {
	preset := auth.LoadPresetFromPath(utils.GetWorkingDirectory(), "autoauth.yml")
	err := preset.IsHeatbeatAlive()
	if err != nil {
		t.Errorf("Heartbeat err: %v", err)
	}
}
