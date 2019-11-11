package options

import "AutoAuthen/pkg/auth"

// KMITL ...
type KMITL struct {
	auth.LoginBase
}

// Login ...
func (req KMITL) Login() error {
	return nil
}
