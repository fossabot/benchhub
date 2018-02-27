package meta

import (
	"sort"
	"sync"

	"github.com/dyweb/gommon/errors"

	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
)

var (
	providerMu        sync.Mutex
	providerFactories = make(map[string]ProviderFactory)
)

type NodeProvider interface {
	// read
	NumNodes() (int, error)
	// TODO: special error for not found?
	// NOTE: we always return by value to avoid (my) common mistake of pointer pointing to last element in for .. range
	FindNodeById(id string) (pbc.Node, error)
	ListNodes() ([]pbc.Node, error)

	// write
	AddNode(id string, node pbc.Node) error
	UpdateNode(id string, node pbc.Node) error

	// delete
	RemoveNode(id string) error
}

type Provider interface {
	NodeProvider
}

// TODO: factory should accept config, this is needed for rdbms
type ProviderFactory func() Provider

func GetProvider(name string) (Provider, error) {
	providerMu.Lock()
	defer providerMu.Unlock()
	if f, ok := providerFactories[name]; !ok {
		return nil, errors.Errorf("provider %s is not registered", name)
	} else {
		// TODO: factory may also return error once we have config
		return f(), nil
	}
}

func RegisterProviderFactory(name string, factory ProviderFactory) {
	providerMu.Lock()
	defer providerMu.Unlock()
	if _, dup := providerFactories[name]; dup {
		log.Panicf("RegisterProviderFactory is called twice for %s", name)
	}
	providerFactories[name] = factory
	// FIXED: this logger never showed up ... because this function is called before the cli application set the
	log.Debugf("register provider factory %s", name)
}

func Providers() []string {
	providerMu.Lock()
	defer providerMu.Unlock()
	list := make([]string, 0, len(providerFactories))
	for name := range providerFactories {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}