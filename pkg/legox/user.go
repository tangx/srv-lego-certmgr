package legox

import (
	"crypto"

	"github.com/go-acme/lego/v4/registration"
)

// You'll need a user or account type that implements acme.user
type user struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *user) GetEmail() string {
	return u.Email
}
func (u user) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *user) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

func (u *user) initial() {
	if u.key == nil {
		u.key = privatekey()
	}
}

func (u *user) register(reg *registration.Resource) {
	u.Registration = reg
}
