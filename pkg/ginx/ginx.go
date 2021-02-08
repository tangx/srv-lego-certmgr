package ginx

import "github.com/gin-gonic/gin"

type Ginx struct {
	server *gin.Engine
	Listen string
}

func NewDefaultServer() *Ginx {
	return &Ginx{
		Listen: ":80",
	}
}

func (g *Ginx) Serve() {
	if err := g.server.Run(g.Listen); err != nil {
		panic(err)
	}
}
