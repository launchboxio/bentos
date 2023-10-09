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

func (d *Docker) CreateCluster() {

}

func (d *Docker) DeleteNetwork() {

}
