package x509util

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func ParseCertificate(cert string) (*x509.Certificate, error) {
	data := []byte(cert)

	// https://blog.csdn.net/u011228889/article/details/81480617
	blk, _ := pem.Decode(data)
	if blk == nil {

		return nil, fmt.Errorf("x509, pem decode failed.")
	}

	return x509.ParseCertificate(blk.Bytes)

}
