package storage

import (
	"github.com/tangx/srv-lego-certmgr/pkg/storage/backend"
)

type Manager struct {
	ClassName string `env:""`

	FileSystem *backend.FileSystem
}

func (mgr *Manager) SetDefaults() {
	if mgr.ClassName == "" {
		mgr.ClassName = "filesystem"
	}
}

func (mgr *Manager) Init() {
	switch mgr.ClassName {
	case "filesystem":
		mgr.FileSystem = &backend.FileSystem{}
	}
}

func (mgr *Manager) Storage() Storager {
	switch mgr.ClassName {
	case "filesystem":
		return mgr.FileSystem
	}

	return nil
}
