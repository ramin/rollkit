package node

import (
	"errors"
	"fmt"

	p2p "github.com/cometbft/cometbft/p2p"
	"github.com/libp2p/go-libp2p/core/crypto"
)

var (
	errNilKey             = errors.New("key can't be nil")
	errUnsupportedKeyType = errors.New("unsupported key type")
)

// GetNodeKey creates libp2p private key from Tendermints NodeKey.
func GetNodeKey(nodeKey *p2p.NodeKey) (crypto.PrivKey, error) {
	if nodeKey == nil || nodeKey.PrivKey == nil {
		return nil, errNilKey
	}
	switch nodeKey.PrivKey.Type() {
	case "ed25519":
		privKey, err := crypto.UnmarshalEd25519PrivateKey(nodeKey.PrivKey.Bytes())
		if err != nil {
			return nil, fmt.Errorf("error while node private key: %w", err)
		}
		return privKey, nil
	default:
		return nil, errUnsupportedKeyType
	}
}
