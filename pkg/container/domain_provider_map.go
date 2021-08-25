package container

import (
	"os"
	"strings"
)

type DomainProviderMap struct {
	container map[string]string
}

func NewDomainProviderMap() *DomainProviderMap {
	dpm := &DomainProviderMap{
		container: make(map[string]string),
	}
	dpm.loadEnv()
	return dpm
}

func (dpm *DomainProviderMap) loadEnv() {
	content := os.Getenv("DomainProviderMap")
	dpm.store(content)
}

func (dpm *DomainProviderMap) Append(content string) {
	dpm.store(content)
}

func (dpm *DomainProviderMap) store(content string) {
	for _, kv := range strings.Split(content, ",") {
		parts := strings.Split(kv, "=")
		if len(parts) != 2 || len(parts[1]) == 0 {
			continue
		}

		dpm.container[parts[0]] = parts[1]
	}
}

func (dpm *DomainProviderMap) Get() map[string]string {
	return dpm.container
}
