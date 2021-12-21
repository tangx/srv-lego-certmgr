package providermanager

import (
	"fmt"

	"github.com/tangx/srv-lego-certmgr/pkg/legox"
)

func NewManager() *Manager {
	return &Manager{
		bucket: make(map[string]legox.Provider),
	}
}

type Manager struct {
	bucket map[string]legox.Provider
}

func (mgr *Manager) Init() {
	if mgr.bucket == nil {
		mgr.bucket = make(map[string]legox.Provider)
	}
}

func (mgr *Manager) Get(name string) legox.Provider {
	p, ok := mgr.bucket[name]
	if ok {
		return p
	}

	return nil
}

func (mgr *Manager) has(name string) bool {
	_, ok := mgr.bucket[name]
	return ok
}

func (mgr *Manager) Register(name string, p legox.Provider) {
	if mgr.has(name) {
		err := fmt.Errorf("%s is already registed", name)
		panic(err)
	}

	mgr.bucket[name] = p
}
