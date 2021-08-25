package cert

import (
	"github.com/gin-gonic/gin"
)

func AppendRoute(root *gin.RouterGroup) {
	rg := root.Group("query")
	{
		rg.GET("/:domain", GetHandler)
		rg.GET("/:domain/download", DownloadHandler)
	}

	root.GET("/list", ListHanlder)
	root.GET("/list-all", ListAllHanlder)
}
