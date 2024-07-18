package gosm

import (
	"context"

	wasm_types "github.com/CosmWasm/wasmd/x/wasm/types"

	rpc_http "github.com/cometbft/cometbft/rpc/client/http"
	core_types "github.com/cometbft/cometbft/rpc/core/types"
)

//go:generate mockery --name GosmClient --filename mock_gosm_client.go
type GosmClient interface {
	// cometbft
	ABCIQuery(ctx context.Context, path string, req Marshaler) (*core_types.ResultABCIQuery, error)

	// cosmwasm
	QuerySmartContractState(ctx context.Context, address string, query any) (*wasm_types.QuerySmartContractStateResponse, error)
}

type RPCClient struct {
	chainID string
	rpcConn CometClient
}

var _ GosmClient = (*RPCClient)(nil)

func NewRPCClient(chainID string, rpcConn CometClient) *RPCClient {
	return &RPCClient{
		chainID: chainID,
		rpcConn: rpcConn,
	}
}

func Dial(ctx context.Context, rpcEndpoint string) (*RPCClient, error) {
	rpcConn, err := rpc_http.New(rpcEndpoint, "/websocket")
	if err != nil {
		return nil, err
	}

	status, err := rpcConn.Status(ctx)
	if err != nil {
		return nil, err
	}

	client := NewRPCClient(status.NodeInfo.Network, rpcConn)

	return client, nil
}
