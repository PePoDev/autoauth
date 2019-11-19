package auth

import (
	"time"

	"github.com/pepodev/xlog"
	"github.com/spf13/viper"
)

// LoadPresetFromPath will load preset from path and return it
func LoadPresetFromPath(dir string, fileName string) AutoAuthPreset {
	viper.SetConfigFile(fileName)
	viper.AddConfigPath(dir)
	if err := viper.ReadInConfig(); err != nil {
		xlog.Fatalf("fatal error config file: %s \n", err)
	}

	basePreset := MainAutoAuthPreset{}
	if err := viper.Unmarshal(&basePreset); err != nil {
		xlog.Fatalf("%v", err)
	}

	return basePreset.AutoAuth
}

// MainAutoAuthPreset is used to map to preset file like yml or json file
type MainAutoAuthPreset struct {
	AutoAuth AutoAuthPreset `mapstructure:"autoauth"`
}

// AutoAuthPreset is base struct contain all configuration of preset file
type AutoAuthPreset struct {
	AutoAuthData `mapstructure:",squash"`

	Login     AutoAuthLogin
	Logout    AutoAuthLogout
	Heartbeat AutoAuthHeartbeat
}

// AutoAuthData contain data for AutoAuth struct
type AutoAuthData struct {
	Name      string
	Encrypted bool

	Save []string

	IsRunning bool
	Try       int
}

// AutoAuthLogin contain login preset
type AutoAuthLogin struct {
	AutoAuthBaseReuqest `mapstructure:",squash"`
}

// AutoAuthLogout contain logout preset
type AutoAuthLogout struct {
	AutoAuthBaseReuqest `mapstructure:",squash"`
}

// AutoAuthHeartbeat contain heartbeat preset
type AutoAuthHeartbeat struct {
	AutoAuthBaseReuqest `mapstructure:",squash"`

	Interval time.Duration
	Retry    int
}

// AutoAuthBaseReuqest is base struct contain data to create request
type AutoAuthBaseReuqest struct {
	Endpoint string
	Method   string
	Header   []string
	Body     []string
	Timeout  time.Duration
}
