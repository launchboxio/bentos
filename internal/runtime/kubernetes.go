package runtime

import (
	"github.com/launchboxio/bentos/internal/runtime/types"
)

type Kubernetes struct {
}

func (k *Kubernetes) CreateNetwork(network *types.CreateNetworkOptions) (*types.CreateNetworkResponse, error) {
	return &types.CreateNetworkResponse{}, nil
}

func (k *Kubernetes) CreateInstance() {

}

func (k *Kubernetes) DeleteNetwork() {

}

func (k *Kubernetes) ListNetworks() (types.NetworkList, error) {
	return types.NetworkList{}, nil
}

func (k *Kubernetes) LoadMetadata() (types.Metadata, error) {
	return types.Metadata{}, nil
}

func (k *Kubernetes) GetNetwork() (types.Network, error) {
	return types.Network{}, nil
}

func (k *Kubernetes) ListInstances() (types.InstanceList, error) {
	return types.InstanceList{}, nil
}
