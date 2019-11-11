package auth

import "fmt"

// LoginBase ...
type LoginBase struct {
}

// Login ...
func (login LoginBase) Login() error {
	return fmt.Errorf("%s not implement %s method", login.ServiceName(), "login")
}

// Logout ...
func (login LoginBase) Logout() error {
	return fmt.Errorf("%s not implement %s method", login.ServiceName(), "logout")
}

// ServiceName ...
func (login LoginBase) ServiceName() string {
	return "[base]"
}

// LoginOption is interface to contain login information
type LoginOption interface {
	ServiceName() string
	Login() error
	Logout() error
}
