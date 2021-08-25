package provider

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/pkg/httpresponse"
)

var dpm = global.DPmapping

func AppendRoute(root *gin.RouterGroup) {
	rg := root.Group("/provider")
	rg.GET("/map", getDpmHandler)
	rg.POST("/map", appendDpmHandler)
}

func getDpmHandler(c *gin.Context) {

	r := dpm.Get()

	httpresponse.StatusOK(c, r)
}

type AppendDP struct {
	Content string `query:"content"`
}

func appendDpmHandler(c *gin.Context) {
	params := &AppendDP{}
	err := ginbinder.ShouldBindRequest(c, params)
	if err != nil {
		httpresponse.StatusBadRequest(c, err)
		return
	}

	dpm.Append(params.Content)
	httpresponse.StatusOK(c, "OK")

}
