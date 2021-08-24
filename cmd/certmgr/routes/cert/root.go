package cert

import (
	"github.com/gin-gonic/gin"
)

var (
	root *gin.RouterGroup
)

func AppendRoute(rg *gin.RouterGroup) {
	root = rg.Group("query")
	{
		root.GET("/:domain", GetHandler)
		root.GET("/:domain/download", DownloadHandler)
	}

	rg.GET("/list", ListHanlder)
	rg.GET("/list-all", ListAllHanlder)
}
