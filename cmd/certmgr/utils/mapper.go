package utils

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-lego-certmgr/cmd/certmgr/global"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
	"github.com/tangx/srv-lego-certmgr/pkg/storage"
)

func newCertMapper() *certMapper {
	m := &certMapper{
		bucket: make(map[string]legox.Certificate),
	}

	m.loading()

	return m
}

type certMapper struct {
	lock   sync.RWMutex
	bucket map[string]legox.Certificate
}

func (m *certMapper) Store(name string, cert legox.Certificate) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.bucket[name] = cert

}

func (m *certMapper) Get(domain string) (legox.Certificate, bool) {

	m.lock.RLock()
	defer m.lock.RUnlock()

	cert, ok := m.bucket[domain]
	if ok {
		return cert, true
	}

	if storager != nil {
		crt := storager.GetByName(domain)
		if crt != nil {
			m.bucket[domain] = *crt

			return *crt, true
		}
	}

	return legox.Certificate{}, false
}

func (m *certMapper) ListValidCerts() map[string]time.Time {
	n := make(map[string]time.Time)
	for _, cert := range m.bucket {
		n[cert.Domain] = cert.NotAfter
	}

	return n
}

func (m *certMapper) ListAll() map[string]legox.Certificate {
	m.loading()

	return m.bucket
}

func (m *certMapper) loading() {

	if storager == nil {
		return
	}

	lister, ok := storager.(storage.StoragerLister)
	if !ok {
		return
	}

	for _, crt := range lister.GetAllCerts() {
		if m.bucket == nil {
			logrus.Panic("m.set is nill")
		}
		if crt == nil {
			continue
		}

		m.Store(crt.Domain, *crt)
	}

	logrus.Debugf("loading from storage done: %s", m.bucket)

}

var (
	mapper   *certMapper
	storager storage.Storager
)

func init() {
	// storager = InitialStorager()
	storager = global.BackendStorage()
	mapper = newCertMapper()

}

// // initial backend storage
// func InitialStorager() storage.Storager {
// 	var sto storage.Storager

// 	viper.AutomaticEnv()

// 	class := viper.GetString("BACKEND_STORAGE_CLASS")
// 	switch class {
// 	case "filesystem":
// 		logrus.Error("create filesystem")
// 		dir := viper.GetString("BACKEND_FILE_SYSTEM_DIR")
// 		sto = backend.NewFileSystem(dir)
// 	}

// 	if sto == nil {
// 		logrus.Error("storage is nill")
// 	}
// 	return sto
// }
