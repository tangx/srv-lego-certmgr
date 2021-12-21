package utils

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
	"github.com/tangx/srv-lego-certmgr/pkg/storage"
	"github.com/tangx/srv-lego-certmgr/pkg/storage/filesystem"
)

func newCertMapper() *certMapper {
	m := &certMapper{
		set: make(map[string]legox.Certificate),
	}

	// m.set["123123"] = legox.Certificate{}
	// fmt.Println("m===>", m)
	m.loading()

	return m
}

type certMapper struct {
	lock sync.RWMutex
	set  map[string]legox.Certificate
	// sto  storage.Storager
}

func (m *certMapper) Store(name string, cert legox.Certificate) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.set[name] = cert

}

func (m *certMapper) Get(domain string) (legox.Certificate, bool) {

	m.lock.RLock()
	defer m.lock.RUnlock()

	cert, ok := m.set[domain]
	if ok {
		return cert, true
	}

	if storager != nil {
		crt := storager.GetByName(domain)
		if crt != nil {
			m.set[domain] = *crt

			return *crt, true
		}
	}

	return legox.Certificate{}, false
}

func (m *certMapper) ListValidCerts() map[string]time.Time {
	n := make(map[string]time.Time)
	for _, cert := range m.set {
		n[cert.Domain] = cert.NotAfter
	}

	return n
}

func (m *certMapper) ListAll() map[string]legox.Certificate {
	m.loading()

	return m.set
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
		if m.set == nil {
			logrus.Panic("m.set is nill")
		}

		// logrus.Error("crt is cert", crt)
		// logrus.Error("m.Store", m.set)

		if crt == nil {
			continue
		}

		m.Store(crt.Domain, *crt)
	}

}

var (
	mapper   *certMapper
	storager storage.Storager
)

func init() {
	storager = InitialStorager()
	mapper = newCertMapper()

}

// initial backend storage
func InitialStorager() storage.Storager {
	var sto storage.Storager

	viper.AutomaticEnv()

	class := viper.GetString("BACKEND_STORAGE_CLASS")
	switch class {
	case "filesystem":
		logrus.Error("create filesystem")
		dir := viper.GetString("BACKEND_FILE_SYSTEM_DIR")
		sto = filesystem.NewStorager(dir)
	}

	if sto == nil {
		logrus.Error("storage is nill")
	}
	return sto
}
