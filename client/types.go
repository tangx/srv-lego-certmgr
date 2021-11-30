package client

import "github.com/tangx/srv-lego-certmgr/pkg/legox"

type GetCertResponse struct {
	Code  int
	Error string
	Data  legox.Certificate
}

type CreateCertReponse struct {
	Code  int
	Error string
	Data  string
}
