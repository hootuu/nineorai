package keys

import (
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"filippo.io/edwards25519"
	"fmt"
	"github.com/mr-tron/base58"
)

type PrivateKey []byte

func NewRandomPrivateKey() (PrivateKey, error) {
	_, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return PrivateKey(pri), nil
}

func MustPrivateKeyFromBase58(in string) PrivateKey {
	out, err := PrivateKeyFromBase58(in)
	if err != nil {
		panic(err)
	}
	return out
}

func PrivateKeyFromBase58(priKey string) (PrivateKey, error) {
	res, err := base58.Decode(priKey)
	if err != nil {
		return nil, err
	}
	// validate
	if _, err := ValidatePrivateKey(res); err != nil {
		return nil, err
	}
	return res, nil
}

func ValidatePrivateKey(b []byte) (bool, error) {
	if len(b) != ed25519.PrivateKeySize {
		return false, fmt.Errorf("invalid private key size, expected %v, got %d", ed25519.PrivateKeySize, len(b))
	}
	// check if the public key is on the ed25519 curve
	pub := ed25519.PrivateKey(b).Public().(ed25519.PublicKey)
	if !IsOnCurve(pub) {
		return false, errors.New("the corresponding public key is NOT on the ed25519 curve")
	}
	return true, nil
}

func IsOnCurve(b []byte) bool {
	if len(b) != ed25519.PublicKeySize {
		return false
	}
	_, err := new(edwards25519.Point).SetBytes(b)
	isOnCurve := err == nil
	return isOnCurve
}

func (k PrivateKey) String() string {
	return base58.Encode(k)
}
