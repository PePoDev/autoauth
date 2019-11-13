package utils

import (
	"time"

	"github.com/pepodev/xlog"

	"github.com/spf13/viper"
)

// ViperGetString ...
func ViperGetString(key string) string {
	if !viper.IsSet(key) {
		xlog.Fatalf("Field %v is empty and doesn't have default value", key)
	}
	return viper.GetString(key)
}

// ViperGetDuration ...
func ViperGetDuration(key string) time.Duration {
	if !viper.IsSet(key) {
		xlog.Fatalf("Field %v is empty and doesn't have default value", key)
	}
	return viper.GetDuration(key)
}

// ViperGetBool ...
func ViperGetBool(key string) bool {
	if !viper.IsSet(key) {
		xlog.Fatalf("Field %v is empty and doesn't have default value", key)
	}
	return viper.GetBool(key)
}

// ViperGetInt ...
func ViperGetInt(key string) int {
	if !viper.IsSet(key) {
		xlog.Fatalf("Field %v is empty and doesn't have default value", key)
	}
	return viper.GetInt(key)
}
