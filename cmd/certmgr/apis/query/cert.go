package query

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/utils"
	"github.com/tangx/srv-lego-certmgr/pkg/httpresponse"
)

func init() {
	QueryRouterGroup.Register(&GetCertByName{})
}

type GetCertByName struct {
	httpx.MethodGet `path:"/:domain"`
	Domain          string `uri:"domain"`
}

func (req *GetCertByName) Output(c *gin.Context) (interface{}, error) {

	cert, ok := utils.GetCert(req.Domain)
	if ok {
		return httpresponse.RespOK(cert), nil
	}

	return httpresponse.RespNotFound(
		errors.New("domain not found"),
	), nil
}
