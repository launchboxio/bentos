package runtime

import (
	"github.com/launchboxio/bentos/internal/runtime/types"
)

type Containerd struct {
}

func (crd *Containerd) CreateNetwork(network *types.CreateNetworkOptions) (*types.CreateNetworkResponse, error) {
	return &types.CreateNetworkResponse{}, nil
}

func (crd *Containerd) CreateInstance() {

}

func (crd *Containerd) DeleteNetwork() {

}

func (crd *Containerd) ListNetworks() (types.NetworkList, error) {
	return types.NetworkList{}, nil
}

func (crd *Containerd) LoadMetadata() (types.Metadata, error) {
	return types.Metadata{}, nil
}

func (crd *Containerd) GetNetwork() (types.Network, error) {
	return types.Network{}, nil
}

func (crd *Containerd) ListInstances() (types.InstanceList, error) {
	return types.InstanceList{}, nil
}
