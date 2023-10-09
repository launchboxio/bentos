package runtime

import (
	"github.com/launchboxio/bentos/internal/runtime/types"
)

// Specify the contract that any runtime adapters
// need to satisfy to be used with BentOS
type Adapter interface {
	CreateNetwork(options *types.CreateNetworkOptions) (*types.CreateNetworkResponse, error)
}
