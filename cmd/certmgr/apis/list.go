package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/utils"
	"github.com/tangx/srv-lego-certmgr/pkg/httpresponse"
)

func init() {
	BaseRouterGroup.Register(&ListCerts{})
	BaseRouterGroup.Register(&ListALLCerts{})
}

type ListCerts struct {
	httpx.MethodGet `path:"/list"`
}

func (req *ListCerts) Output(c *gin.Context) (interface{}, error) {
	r := utils.ListValidCerts()
	ret := httpresponse.RespOK(r)

	return ret, nil
}

type ListALLCerts struct {
	httpx.MethodGet `path:"/list-all"`
}

func (req *ListALLCerts) Output(c *gin.Context) (interface{}, error) {
	r := utils.ListAll()
	ret := httpresponse.RespOK(r)

	return ret, nil
}
