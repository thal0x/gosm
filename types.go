package gosm

import (
	"context"

	"github.com/cometbft/cometbft/rpc/client"
)

//go:generate mockery --name CometClient --filename mock_comet_client.go
type CometClient interface {
	client.RemoteClient
}

type Marshaler interface {
	Marshal() ([]byte, error)
}

//go:generate mockery --name Caller --filename mock_caller.go
type Caller interface {
	Call(ctx context.Context, method string, params map[string]interface{}, result interface{}) (interface{}, error)
}
