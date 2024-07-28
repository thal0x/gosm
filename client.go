package gosm

import (
	"context"

	"golang.org/x/mod/semver"

	wasm_types "github.com/CosmWasm/wasmd/x/wasm/types"

	rpc_http "github.com/cometbft/cometbft/rpc/client/http"
	core_types "github.com/cometbft/cometbft/rpc/core/types"
	jsonrpc_client "github.com/cometbft/cometbft/rpc/jsonrpc/client"
)

const (
	cometEncodingThreshold = "v0.37.0-alpha"
)

//go:generate mockery --name GosmClient --filename mock_gosm_client.go
type GosmClient interface {
	// cometbft
	ABCIQuery(ctx context.Context, path string, req Marshaler) (*core_types.ResultABCIQuery, error)
	Block(ctx context.Context, height *int64) (*core_types.ResultBlock, error)
	BlockResults(ctx context.Context, height *int64) (*core_types.ResultBlockResults, error)
	CometLegacyEncoding() bool

	// cosmwasm
	QuerySmartContractState(ctx context.Context, address string, query any) (*wasm_types.QuerySmartContractStateResponse, error)
}

type RPCClient struct {
	chainID             string
	rpcConn             CometClient
	rpcCaller           jsonrpc_client.Caller
	cometLegacyEncoding bool
}

var _ GosmClient = (*RPCClient)(nil)

func NewRPCClient(
	chainID string,
	rpcConn CometClient,
	rpcCaller jsonrpc_client.Caller,
	cometLegacyEncoding bool,
) *RPCClient {
	return &RPCClient{
		chainID:             chainID,
		rpcConn:             rpcConn,
		rpcCaller:           rpcCaller,
		cometLegacyEncoding: cometLegacyEncoding,
	}
}

func Dial(ctx context.Context, rpcEndpoint string) (*RPCClient, error) {
	rpcConn, err := rpc_http.New(rpcEndpoint, "/websocket")
	if err != nil {
		return nil, err
	}

	rpcCaller, err := jsonrpc_client.New(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	status, err := rpcConn.Status(ctx)
	if err != nil {
		return nil, err
	}

	client := NewRPCClient(status.NodeInfo.Network, rpcConn, rpcCaller, useLegacyEncodedEvents(status.NodeInfo.Version))

	return client, nil
}

func useLegacyEncodedEvents(version string) bool {
	return semver.Compare("v"+version, cometEncodingThreshold) < 0
}
