package gosm

import (
	"context"
	"fmt"

	core_types "github.com/cometbft/cometbft/rpc/core/types"
	comet_types "github.com/cometbft/cometbft/types"
)

func (c *RPCClient) CometLegacyEncoding() bool {
	return c.cometLegacyEncoding
}

func (c *RPCClient) Block(ctx context.Context, height *int64) (*core_types.ResultBlock, error) {
	if c.chainID == "pacific-1" {
		return c.blockSei(ctx, height)
	}

	return c.rpcConn.Block(ctx, height)
}

func (c *RPCClient) blockSei(ctx context.Context, height *int64) (*core_types.ResultBlock, error) {
	result := new(SeiResultBlock)

	params := make(map[string]interface{})
	params["height"] = height

	_, err := c.rpcCaller.Call(ctx, "block", params, result)
	if err != nil {
		fmt.Println(result.BlockID)
		return nil, err
	}

	return &core_types.ResultBlock{
		BlockID: comet_types.BlockID{
			Hash:          result.BlockID.Hash,
			PartSetHeader: result.BlockID.PartSetHeader,
		},
		Block: &comet_types.Block{
			Header:     result.Block.Header,
			Data:       result.Block.Data,
			LastCommit: result.Block.LastCommit,
		},
	}, nil
}

type SeiResultBlock struct {
	BlockID *comet_types.BlockID `json:"block_id"`
	Block   *SeiBlock            `json:"block"`
}

type SeiBlock struct {
	Header comet_types.Header `json:"header"`
	Data   comet_types.Data   `json:"data"`
	// Evidence comet_types.EvidenceData `json:"evidence"`
	LastCommit *comet_types.Commit `json:"last_commit"`
}

func (c *RPCClient) BlockResults(ctx context.Context, height *int64) (*core_types.ResultBlockResults, error) {
	return c.rpcConn.BlockResults(ctx, height)
}
