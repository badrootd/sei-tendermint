package state

import (
	tmstate "github.com/badrootd/sei-tendermint/proto/tendermint/state"
	"github.com/badrootd/sei-tendermint/types"
)

func ABCIResponsesResultsHash(ar *tmstate.ABCIResponses) []byte {
	return types.NewResults(ar.DeliverTxs).Hash()
}
