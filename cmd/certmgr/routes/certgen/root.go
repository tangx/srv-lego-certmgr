package certgen

import "github.com/gin-gonic/gin"

// 设置路由
func AppendRoute(rg *gin.RouterGroup, name string) {

	subroot := rg.Group("gen")

	provroot := subroot.Group(name)
	// provroot.GET("/ping")
	provroot.POST("/:domains", ApplyCertificateHandler)
	provroot.GET("/:domains", GetHandler)

}
