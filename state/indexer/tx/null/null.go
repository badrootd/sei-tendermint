package null

import (
	"context"
	"errors"
	"github.com/badrootd/sei-tendermint/state/indexer"

	abci "github.com/badrootd/sei-tendermint/abci/types"
	"github.com/badrootd/sei-tendermint/internal/pubsub/query"
)

var _ indexer.TxIndexer = (*TxIndex)(nil)

// TxIndex acts as a /dev/null.
type TxIndex struct{}

// Get on a TxIndex is disabled and panics when invoked.
func (txi *TxIndex) Get(hash []byte) (*abci.TxResult, error) {
	return nil, errors.New(`indexing is disabled (set 'tx_index = "kv"' in config)`)
}

// AddBatch is a noop and always returns nil.
func (txi *TxIndex) AddBatch(batch *indexer.Batch) error {
	return nil
}

// Index is a noop and always returns nil.
func (txi *TxIndex) Index(results []*abci.TxResult) error {
	return nil
}

func (txi *TxIndex) Search(ctx context.Context, q *query.Query) ([]*abci.TxResult, error) {
	return []*abci.TxResult{}, nil
}
