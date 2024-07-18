package gosm

import (
	"context"

	core_types "github.com/cometbft/cometbft/rpc/core/types"
)

func (client *RPCClient) ABCIQuery(ctx context.Context, path string, req Marshaler) (*core_types.ResultABCIQuery, error) {
	data, err := req.Marshal()
	if err != nil {
		return nil, err
	}

	result, err := client.rpcConn.ABCIQuery(ctx, path, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
