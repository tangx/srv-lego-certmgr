package gen

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/rum-gonic/rum"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
)

var GenRouterGroup = rum.NewRouterGroup("/gen")

type GenerateCertByDomain struct {
	httpx.MethodPost `path:"/:domain"`
	Domain           string `uri:"domain"`
}

func (req *GenerateCertByDomain) Output(c *gin.Context) (interface{}, error) {

	return nil, nil
}

func init() {
	for _, name := range global.Providers() {
		rg := rum.NewRouterGroup(name)
		GenRouterGroup.Register(rg)

		rg.Register(&GenerateCertByDomain{})
	}
}
