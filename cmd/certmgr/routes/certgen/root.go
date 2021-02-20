package certgen

import "github.com/gin-gonic/gin"

// 设置路由
func AppendRoute(rg *gin.RouterGroup, name string) {

	subroot := rg.Group("gen")

	// 根据 provider 注册路由
	provroot := subroot.Group(name)
	provroot.POST("/:domains", ApplyCertificateHandler)
	provroot.GET("/:domains", GetHandler)
	// 启动重试队列
	retryApply(name)
}
