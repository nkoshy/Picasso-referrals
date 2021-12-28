package codec

import (
	"github.com/PicassoExchange/picasso-referals/src/ethereum/crypto/ethsecp256k1"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

var (
	_ cryptotypes.PubKey  = &ethsecp256k1.PubKey{}
	_ cryptotypes.PrivKey = &ethsecp256k1.PrivKey{}
)

// RegisterInterfaces registers the ethsecp256k1 implementations of tendermint crypto types.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PrivKey)(nil), &ethsecp256k1.PrivKey{})
}
