package gosm

import "github.com/cometbft/cometbft/rpc/client"

//go:generate mockery --name CometClient --filename mock_comet_client.go
type CometClient interface {
	client.RemoteClient
}
