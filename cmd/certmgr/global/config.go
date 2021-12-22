package global

import (
	"github.com/go-jarvis/jarvis"
	"github.com/go-jarvis/jarvis/pkg/appctx"
	"github.com/go-jarvis/rum-gonic/confhttp"
	"github.com/tangx/srv-lego-certmgr/pkg/legox/alidnsprovider"
	"github.com/tangx/srv-lego-certmgr/pkg/legox/dnspodprovider"
	providermanager "github.com/tangx/srv-lego-certmgr/pkg/legox/provider-manager"
	"github.com/tangx/srv-lego-certmgr/pkg/storage"
)

var (
	// Server  = ginx.Default()
	Appname = "LegoCertManager"
)

// config fields
var (
	server         = &confhttp.Server{}
	alidns         = &alidnsprovider.Config{}
	dnspod         = &dnspodprovider.Config{}
	storageManager = &storage.Manager{}
)

var (
	providers = providermanager.NewManager()

	App = jarvis.New().WithOptions(
		appctx.WithName(Appname),
		appctx.WithRoot("../.."),
	)
)

func init() {
	config := &struct {
		HttpServer     *confhttp.Server
		Alidns         *alidnsprovider.Config
		Dnspod         *dnspodprovider.Config
		BackendManager *storage.Manager
	}{
		HttpServer:     server,
		Alidns:         alidns,
		Dnspod:         dnspod,
		BackendManager: storageManager,
	}

	_ = App.Conf(config)

	registerProviders()
}

func Server() *confhttp.Server {
	return server
}

func registerProviders() {

	if alidns.Enabled {
		providers.Register(alidns.NickName(), alidns)
	}

	if dnspod.Enabled {
		providers.Register(dnspod.NickName(), dnspod)
	}
}

func ProviderManager() *providermanager.Manager {
	return providers
}

func BackendStorage() storage.Storager {
	return storageManager.Storage()
}
