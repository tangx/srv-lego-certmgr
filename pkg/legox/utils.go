package legox

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/sirupsen/logrus"
)

func privatekey() *ecdsa.PrivateKey {
	// Create a user. New accounts need an email and private key to start.
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		logrus.Fatal(err)
	}

	return key
}
