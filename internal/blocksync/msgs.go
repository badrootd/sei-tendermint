package blocksync

import (
	bcproto "github.com/badrootd/sei-tendermint/proto/tendermint/blocksync"
	"github.com/badrootd/sei-tendermint/types"
)

const (
	MaxMsgSize = types.MaxBlockSizeBytes +
		bcproto.BlockResponseMessagePrefixSize +
		bcproto.BlockResponseMessageFieldKeySize
)
