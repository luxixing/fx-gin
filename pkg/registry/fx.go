package registry

import (
	"sync"

	"go.uber.org/fx"
)

var (
	modules []fx.Option
	mu      sync.Mutex
)

// Register adds a new fx option to the registry
func Register(opt fx.Option) {
	mu.Lock()
	defer mu.Unlock()
	modules = append(modules, opt)
}

// GetModules returns all registered modules
func GetModules() fx.Option {
	mu.Lock()
	defer mu.Unlock()
	return fx.Options(modules...)
}
