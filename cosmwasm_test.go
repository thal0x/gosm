package gosm

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	wasm_types "github.com/CosmWasm/wasmd/x/wasm/types"
	abci_types "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/bytes"
	core_types "github.com/cometbft/cometbft/rpc/core/types"

	"github.com/thal0x/gosm/mocks"
	"github.com/thal0x/gosm/utils"
)

func TestQuerySmartContractState(t *testing.T) {
	ctx := context.Background()

	mockCometClient := mocks.NewCometClient(t)
	mockCometClient.On("ABCIQuery", ctx, "/cosmwasm.wasm.v1.Query/SmartContractState", bytes.HexBytes(utils.Must(hex.DecodeString("0A426E657574726F6E316A6A7873357239786E666A736564687379396A61686779746A746D72346B6135396D7368727171687A6A37717338376A7778667174716E676767120B7B2270616972223A7B7D7D")))).Return(&core_types.ResultABCIQuery{
		Response: abci_types.ResponseQuery{
			Code:  0,
			Value: utils.Must(base64.StdEncoding.DecodeString("CugCeyJhc3NldF9pbmZvcyI6W3sibmF0aXZlX3Rva2VuIjp7ImRlbm9tIjoiaWJjL0YwODJCNjVDODhFNEI2RDVFRjFEQjI0M0NEQTFEMzMxRDAwMjc1OUU5MzhBMEY1Q0QzRkZEQzVENTNCM0UzNDkifX0seyJuYXRpdmVfdG9rZW4iOnsiZGVub20iOiJ1bnRybiJ9fV0sImNvbnRyYWN0X2FkZHIiOiJuZXV0cm9uMWpqeHM1cjl4bmZqc2VkaHN5OWphaGd5dGp0bXI0a2E1OW1zaHJxcWh6ajdxczg3and4ZnF0cW5nZ2ciLCJsaXF1aWRpdHlfdG9rZW4iOiJuZXV0cm9uMTM3cnprYWNyeXpzdGZ5eWx2dnZ1MnZxOXVqNWw4OXl6eDl2N2d4cDBzNzd4cTVhY2gyeHMzdDZ0M3AiLCJwYWlyX3R5cGUiOnsiY3VzdG9tIjoiY29uY2VudHJhdGVkIn19")),
		},
	}, nil)

	client := &RPCClient{
		chainID: "test-chain",
		rpcConn: mockCometClient,
	}

	type QueryMsg struct {
		Pair struct{} `json:"pair"`
	}

	result, err := client.QuerySmartContractState(ctx, "neutron1jjxs5r9xnfjsedhsy9jahgytjtmr4ka59mshrqqhzj7qs87jwxfqtqnggg", &QueryMsg{
		Pair: struct{}{},
	})
	assert.NoError(t, err)

	expected := &wasm_types.QuerySmartContractStateResponse{
		Data: []byte(`{"asset_infos":[{"native_token":{"denom":"ibc/F082B65C88E4B6D5EF1DB243CDA1D331D002759E938A0F5CD3FFDC5D53B3E349"}},{"native_token":{"denom":"untrn"}}],"contract_addr":"neutron1jjxs5r9xnfjsedhsy9jahgytjtmr4ka59mshrqqhzj7qs87jwxfqtqnggg","liquidity_token":"neutron137rzkacryzstfyylvvvu2vq9uj5l89yzx9v7gxp0s77xq5ach2xs3t6t3p","pair_type":{"custom":"concentrated"}}`),
	}

	assert.Equal(t, expected, result)
}
