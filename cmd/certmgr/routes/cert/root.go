package cert

import (
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	root *gin.RouterGroup
)

func AppendRoute(rg *gin.RouterGroup) {
	root = rg.Group("cert")

	query := root.Group("query")
	{
		query.GET("/:domain", GetHandler)
		query.GET("/:domain/download", DownloadHandler)
	}

	root.GET("/list", ListHanlder)
}

func convert(domain string) string {

	l := strings.Split(domain, ".")
	if l[0] == "wild" {
		l[0] = "*"
	}

	return strings.Join(l, ".")
}
