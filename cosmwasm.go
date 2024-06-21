package gosm

import (
	"context"
	"encoding/json"
	"fmt"

	wasm_types "github.com/CosmWasm/wasmd/x/wasm/types"
)

func (client *RPCClient) QuerySmartContractState(ctx context.Context, address string, query any) (*wasm_types.QuerySmartContractStateResponse, error) {
	data, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	queryMsg := &wasm_types.QuerySmartContractStateRequest{
		Address:   address,
		QueryData: data,
	}

	data, err = queryMsg.Marshal()
	if err != nil {
		return nil, err
	}

	result, err := client.rpcConn.ABCIQuery(ctx, "/cosmwasm.wasm.v1.Query/SmartContractState", data)
	if err != nil {
		return nil, err
	}

	if result.Response.Code != 0 {
		return nil, fmt.Errorf("QuerySmartContractState failed: %s", result.Response.Log)
	}

	var response wasm_types.QuerySmartContractStateResponse
	err = response.Unmarshal(result.Response.Value)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
