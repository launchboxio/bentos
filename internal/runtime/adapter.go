package runtime

import (
	"errors"
	"github.com/launchboxio/bentos/internal/runtime/types"
)

// Specify the contract that any runtime adapters
// need to satisfy to be used with BentOS
type Adapter interface {
	CreateNetwork(options *types.CreateNetworkOptions) (*types.CreateNetworkResponse, error)
	ListNetworks() (types.NetworkList, error)
	GetNetwork() (types.Network, error)
	LoadMetadata() (types.Metadata, error)
	ListInstances() (types.InstanceList, error)
}

func Load(runtime types.Runtime) (Adapter, error) {
	switch runtime {
	case types.Docker:
		return &Docker{}, nil
	case types.KVM:
		return &KVM{}, nil
	case types.Xen:
		return &Xen{}, nil
	case types.Containerd:
		return &Containerd{}, nil
	case types.Kubernetes:
		return &Kubernetes{}, nil
	}
	return nil, errors.New("Unsupported runtime: " + runtime.ToString())
}
