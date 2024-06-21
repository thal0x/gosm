package gosm

import (
	"context"

	wasm_types "github.com/CosmWasm/wasmd/x/wasm/types"
	rpc_http "github.com/cometbft/cometbft/rpc/client/http"
)

//go:generate mockery --name GosmClient --filename mock_gosm_client.go
type GosmClient interface {
	// cosmwasm
	QuerySmartContractState(ctx context.Context, address string, query any) (*wasm_types.QuerySmartContractStateResponse, error)
}

type RPCClient struct {
	chainID string
	rpcConn CometClient
}

var _ GosmClient = (*RPCClient)(nil)

func NewRPCClient(ctx context.Context, rpcConn CometClient) (*RPCClient, error) {
	status, err := rpcConn.Status(ctx)
	if err != nil {
		return nil, err
	}

	return &RPCClient{
		chainID: status.NodeInfo.Network,
		rpcConn: rpcConn,
	}, nil
}

func Dial(ctx context.Context, rpcEndpoint string) (*RPCClient, error) {
	rpcConn, err := rpc_http.New(rpcEndpoint, "/websocket")
	if err != nil {
		return nil, err
	}

	return NewRPCClient(ctx, rpcConn)
}
