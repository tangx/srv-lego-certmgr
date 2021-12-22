// 保存在本地目录
package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

// var _ storage.Storager = (*FileSystem)(nil)

func NewFileSystem(dir string) *FileSystem {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return &FileSystem{
		DirPath: dir,
	}
}

type FileSystem struct {
	DirPath string `env:""`
}

func (fs *FileSystem) SetDefaults() {
	if fs.DirPath == "" {
		fs.DirPath = "lego-certmgr"
	}
}

// Store a cert into filesytem
func (fs *FileSystem) Store(cert *legox.Certificate) bool {
	file := fs.abs(cert.Domain)
	data, err := json.Marshal(cert)
	if err != nil {
		return false
	}

	err = ioutil.WriteFile(file, data, os.ModePerm)
	return err == nil
}

// GetByName return a cert object, or nil if error
func (fs *FileSystem) GetByName(name string) *legox.Certificate {
	file := fs.abs(name)
	fmt.Println(file)
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

// GetAllCerts all cert
func (fs *FileSystem) GetAllCerts() []*legox.Certificate {
	all := make([]*legox.Certificate, 0)
	for _, name := range fs.walk() {
		cert := fs.GetByName(name)

		if cert == nil {
			continue
		}

		all = append(all, cert)
	}
	return all
}

// abs return real path of domain
func (fs *FileSystem) abs(name string) string {

	name = filepath.Join(fs.DirPath, name)
	if strings.HasSuffix(name, ".json") {
		return name
	}

	return name + ".json"

}

func (fs *FileSystem) walk() []string {
	files, err := os.ReadDir(fs.DirPath)
	if err != nil {
		logrus.Fatalf("os.ReadDir failed: %v", err)
	}

	all := []string{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		all = append(all, f.Name())

	}

	return all
}
