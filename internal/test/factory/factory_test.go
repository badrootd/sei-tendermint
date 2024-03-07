package factory

import (
	"testing"

	"github.com/badrootd/sei-tendermint/types"
)

func TestMakeHeader(t *testing.T) {
	MakeHeader(t, &types.Header{})
}

func TestRandomNodeID(t *testing.T) {
	RandomNodeID(t)
}
