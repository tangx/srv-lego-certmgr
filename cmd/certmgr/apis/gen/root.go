package gen

import (
	"github.com/go-jarvis/rum-gonic/rum"
)

var GenRouterGroup = rum.NewRouterGroup("/gen")

func init() {
	GenRouterGroup.Register(&GenerateCertByDomain{})
}
