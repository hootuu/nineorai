package keys

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"filippo.io/edwards25519"
	"github.com/hootuu/gelato/crtpto/ed25519x"
	"github.com/hootuu/gelato/errors"
	"github.com/mr-tron/base58"
)

const (
	PublicKeyLength = 32
)

type PrivateKey []byte
type PublicKey [PublicKeyLength]byte

type Key struct {
	Private PrivateKey `bson:"private" json:"private"`
	Public  PublicKey  `bson:"public" json:"public"`
	addr    Address
}

func NewKey() (*Key, *errors.Error) {
	pub, pri, err := ed25519x.NewRandom()
	if err != nil {
		return nil, err
	}
	key := &Key{
		Private: pri,
	}
	copy(key.Public[:], pub)
	return key, nil
}

func (k *Key) Address() Address {
	if len(string(k.addr)) == 0 {
		k.addr = k.Public.Address()
	}
	return k.addr
}

func (k PrivateKey) PublicKey() PublicKey {
	if err := k.Validate(); err != nil {
		panic(err)
	}

	p := ed25519.PrivateKey(k)
	pub := p.Public().(ed25519.PublicKey)

	var publicKey PublicKey
	copy(publicKey[:], pub)

	return publicKey
}

func (k PrivateKey) Validate() *errors.Error {
	_, err := PrivateKeyValidate(k)
	return err
}

func (k PrivateKey) ToBase58() string {
	return base58.Encode(k)
}

func (k PrivateKey) Sign(payload []byte) (Signature, *errors.Error) {
	if err := k.Validate(); err != nil {
		return Signature{}, errors.E("input_invalid", err)
	}
	p := ed25519.PrivateKey(k)
	signData, err := p.Sign(rand.Reader, payload, crypto.Hash(0))
	if err != nil {
		return Signature{}, errors.E("input_invalid", err)
	}

	var signature Signature
	copy(signature[:], signData)

	return signature, nil
}

func (p PublicKey) Verify(message []byte, signature Signature) bool {
	pub := ed25519.PublicKey(p[:])
	return ed25519.Verify(pub, message, signature[:])
}

func (p PublicKey) ToBase58() string {
	return base58.Encode(p[:])
}

func (p PublicKey) Address() Address {
	return Address(p.ToBase58())
}

func PrivateKeyFromBase58(pri string) (PrivateKey, *errors.Error) {
	res, err := base58.Decode(pri)
	if err != nil {
		return nil, errors.E("input_invalid", err)
	}
	// validate
	if _, err := PrivateKeyValidate(res); err != nil {
		return nil, err
	}
	return res, nil
}

func PrivateKeyValidate(b []byte) (bool, *errors.Error) {
	if len(b) != ed25519.PrivateKeySize {
		return false, errors.E("input_invalid",
			"invalid private key size, expected %v, got %d", ed25519.PrivateKeySize, len(b))
	}
	// check if the public key is on the ed25519 curve
	pub := ed25519.PrivateKey(b).Public().(ed25519.PublicKey)
	if !PublicKeyIsOnCurve(pub) {
		return false, errors.E("input_invalid", "the corresponding public key is NOT on the ed25519 curve")
	}
	return true, nil
}

func PublicKeyIsOnCurve(b []byte) bool {
	if len(b) != ed25519.PublicKeySize {
		return false
	}
	_, err := new(edwards25519.Point).SetBytes(b)
	isOnCurve := err == nil
	return isOnCurve
}

func PublicKeyFromBytes(in []byte) (out PublicKey, err *errors.Error) {
	byteCount := len(in)

	if byteCount != PublicKeyLength {
		err = errors.E("input_invalid", "invalid public key size, expected %v, got %d", PublicKeyLength, byteCount)
		return
	}

	copy(out[:], in)
	return
}

func PublicKeyFromBase58(in string) (out PublicKey, err *errors.Error) {
	val, nErr := base58.Decode(in)
	if nErr != nil {
		err = errors.E("input_invalid", "decode: %s", nErr.Error())
		return
	}

	if len(val) != PublicKeyLength {
		err = errors.E("input_invalid", "invalid length, expected %v, got %d", PublicKeyLength, len(val))
		return
	}

	copy(out[:], val)
	return
}
