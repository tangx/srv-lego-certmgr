package storage

import "github.com/tangx/srv-lego-certmgr/pkg/legox"

type Storager interface {
	// 保存一条数据
	Store(*legox.Certificate) bool

	// 读取一条数据
	GetByName(name string) *legox.Certificate
}

type StoragerLister interface {
	GetAllCerts() []*legox.Certificate
}
