package runtime

import "github.com/launchboxio/bentos/internal/runtime/types"

type KVM struct {
}

func (k *KVM) CreateNetwork(network *types.CreateNetworkOptions) (*types.CreateNetworkResponse, error) {
	return &types.CreateNetworkResponse{}, nil
}

func (k *KVM) CreateInstance() {

}

func (k *KVM) DeleteNetwork() {

}

func (k *KVM) ListNetworks() (types.NetworkList, error) {
	return types.NetworkList{}, nil
}

func (k *KVM) LoadMetadata() (types.Metadata, error) {
	return types.Metadata{}, nil
}

func (k *KVM) GetNetwork() (types.Network, error) {
	return types.Network{}, nil
}

func (k *KVM) ListInstances() (types.InstanceList, error) {
	return types.InstanceList{}, nil
}
