package runtime

import (
	"github.com/launchboxio/bentos/internal/runtime/types"
)

type Docker struct {
}

func (d *Docker) CreateNetwork(network *types.CreateNetworkOptions) (*types.CreateNetworkResponse, error) {
	return &types.CreateNetworkResponse{}, nil
}

func (d *Docker) CreateInstance() {

}

func (d *Docker) DeleteNetwork() {

}

func (d *Docker) ListNetworks() (types.NetworkList, error) {
	return types.NetworkList{}, nil
}

func (d *Docker) LoadMetadata() (types.Metadata, error) {
	return types.Metadata{}, nil
}

func (d *Docker) GetNetwork() (types.Network, error) {
	return types.Network{}, nil
}

func (d *Docker) ListInstances() (types.InstanceList, error) {
	return types.InstanceList{}, nil
}
