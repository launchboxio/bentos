package runtime

import (
	"github.com/launchboxio/bentos/internal/runtime/types"
)

type Xen struct {
}

func (x *Xen) CreateNetwork(network *types.CreateNetworkOptions) (*types.CreateNetworkResponse, error) {
	return &types.CreateNetworkResponse{}, nil
}

func (x *Xen) CreateInstance() {

}

func (x *Xen) DeleteNetwork() {

}

func (x *Xen) ListNetworks() (types.NetworkList, error) {
	return types.NetworkList{}, nil
}

func (x *Xen) LoadMetadata() (types.Metadata, error) {
	return types.Metadata{}, nil
}

func (x *Xen) GetNetwork() (types.Network, error) {
	return types.Network{}, nil
}

func (x *Xen) ListInstances() (types.InstanceList, error) {
	return types.InstanceList{}, nil
}
