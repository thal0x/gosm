package gosm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cometbft/cometbft/p2p"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"

	"github.com/thal0x/gosm/mocks"
)

func TestNewRPCClient(t *testing.T) {
	ctx := context.Background()

	mockCometClient := mocks.NewCometClient(t)
	mockCometClient.On("Status", ctx).Return(&ctypes.ResultStatus{
		NodeInfo: p2p.DefaultNodeInfo{
			Network: "chain-1",
		},
	}, nil)

	client, err := NewRPCClient(ctx, mockCometClient)
	assert.NoError(t, err)

	assert.Equal(t, client.chainID, "chain-1")
	assert.Equal(t, client.rpcConn, mockCometClient)
}
