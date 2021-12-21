// 保存在本地目录
package filesystem

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tangx/srv-lego-certmgr/pkg/legox"
	"github.com/tangx/srv-lego-certmgr/pkg/storage"
)

var _ storage.Storager = (*Storager)(nil)

func NewStorager(dir string) *Storager {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return &Storager{
		dir: dir,
	}
}

type Storager struct {
	dir string
}

// Store a cert into filesytem
func (fs *Storager) Store(cert *legox.Certificate) bool {
	file := fs.abs(cert.Domain)
	data, err := json.Marshal(cert)
	if err != nil {
		return false
	}

	err = ioutil.WriteFile(file, data, os.ModePerm)
	return err == nil
}

// GetByName return a cert object, or nil if error
func (fs *Storager) GetByName(name string) *legox.Certificate {
	file := fs.abs(name)
	data, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	cert := &legox.Certificate{}
	err = json.Unmarshal(data, cert)
	if err != nil {
		return nil
	}

	return cert
}

// abs return real path of domain
func (fs *Storager) abs(name string) string {
	return filepath.Join(fs.dir, name, ".txt")
}
