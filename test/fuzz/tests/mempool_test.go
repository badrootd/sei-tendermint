//go:build gofuzz || go1.18

package tests

import (
	"context"
	"testing"

	abciclient "github.com/badrootd/sei-tendermint/abci/client"
	"github.com/badrootd/sei-tendermint/abci/example/kvstore"
	"github.com/badrootd/sei-tendermint/config"
	"github.com/badrootd/sei-tendermint/internal/mempool"
	"github.com/badrootd/sei-tendermint/libs/log"
	"github.com/badrootd/sei-tendermint/types"
)

type TestPeerEvictor struct {
	evicting map[types.NodeID]struct{}
}

func NewTestPeerEvictor() *TestPeerEvictor {
	return &TestPeerEvictor{evicting: map[types.NodeID]struct{}{}}
}

func (e *TestPeerEvictor) Errored(peerID types.NodeID, err error) {
	e.evicting[peerID] = struct{}{}
}

func FuzzMempool(f *testing.F) {
	app := kvstore.NewApplication()
	logger := log.NewNopLogger()
	conn := abciclient.NewLocalClient(logger, app)
	err := conn.Start(context.TODO())
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false

	mp := mempool.NewTxMempool(logger, cfg, conn, NewTestPeerEvictor())

	f.Fuzz(func(t *testing.T, data []byte) {
		_ = mp.CheckTx(context.Background(), data, nil, mempool.TxInfo{})
	})
}
