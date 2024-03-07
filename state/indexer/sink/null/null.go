package null

import (
	"context"
	"github.com/badrootd/sei-tendermint/state/indexer"

	abci "github.com/badrootd/sei-tendermint/abci/types"
	"github.com/badrootd/sei-tendermint/internal/pubsub/query"
	"github.com/badrootd/sei-tendermint/types"
)

var _ indexer.EventSink = (*EventSink)(nil)

// EventSink implements a no-op indexer.
type EventSink struct{}

func NewEventSink() indexer.EventSink {
	return &EventSink{}
}

func (nes *EventSink) Type() indexer.EventSinkType {
	return indexer.NULL
}

func (nes *EventSink) IndexBlockEvents(bh types.EventDataNewBlockHeader) error {
	return nil
}

func (nes *EventSink) IndexTxEvents(results []*abci.TxResult) error {
	return nil
}

func (nes *EventSink) SearchBlockEvents(ctx context.Context, q *query.Query) ([]int64, error) {
	return nil, nil
}

func (nes *EventSink) SearchTxEvents(ctx context.Context, q *query.Query) ([]*abci.TxResult, error) {
	return nil, nil
}

func (nes *EventSink) GetTxByHash(hash []byte) (*abci.TxResult, error) {
	return nil, nil
}

func (nes *EventSink) HasBlock(h int64) (bool, error) {
	return false, nil
}

func (nes *EventSink) Stop() error {
	return nil
}
