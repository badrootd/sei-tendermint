package sr25519

import (
	"github.com/badrootd/sei-tendermint/internal/jsontypes"
	tmjson "github.com/badrootd/sei-tendermint/libs/json"
)

const (
	PrivKeyName = "tendermint/PrivKeySr25519"
	PubKeyName  = "tendermint/PubKeySr25519"
)

func init() {
	tmjson.RegisterType(PubKey{}, PubKeyName)
	tmjson.RegisterType(PrivKey{}, PrivKeyName)

	jsontypes.MustRegister(PubKey{})
	jsontypes.MustRegister(PrivKey{})
}
