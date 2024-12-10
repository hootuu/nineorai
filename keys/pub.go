package keys

import (
	"fmt"
	"github.com/mr-tron/base58"
)

const (
	PublicKeyLength = 32
)

type PublicKey [PublicKeyLength]byte

func PublicKeyFromBytes(in []byte) (out PublicKey) {
	byteCount := len(in)

	if byteCount != PublicKeyLength {
		panic(fmt.Errorf("invalid public key size, expected %v, got %d", PublicKeyLength, byteCount))
	}

	copy(out[:], in)
	return
}

func MustPublicKeyFromBase58(in string) PublicKey {
	out, err := PublicKeyFromBase58(in)
	if err != nil {
		panic(err)
	}
	return out
}

// PublicKeyFromBase58 creates a PublicKey from a base58 encoded string.
// NOTE: it will accept on- and off-curve pubkeys.
func PublicKeyFromBase58(in string) (out PublicKey, err error) {
	val, err := base58.Decode(in)
	if err != nil {
		return out, fmt.Errorf("decode: %w", err)
	}

	if len(val) != PublicKeyLength {
		return out, fmt.Errorf("invalid length, expected %v, got %d", PublicKeyLength, len(val))
	}

	copy(out[:], val)
	return
}

func (p PublicKey) Equals(pb PublicKey) bool {
	return p == pb
}

func (p PublicKey) Bytes() []byte {
	return []byte(p[:])
}

func (p PublicKey) String() string {
	return base58.Encode(p[:])
}
