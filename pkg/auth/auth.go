package auth

import (
	"fmt"
)

// StartAutoLogin ...
func StartAutoLogin() {

}

// StopAutoLogin ...
func StopAutoLogin() {

}

// LoginWithOption will create request to authentication service
func LoginWithOption(option LoginOption) {
	if err := option.Login(); err != nil {
		fmt.Printf("error: %v", err)
	}
}
