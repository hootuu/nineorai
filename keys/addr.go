package keys

import "github.com/hootuu/gelato/errors"

type Address string

func (addr Address) PublicKey() (PublicKey, *errors.Error) {
	return PublicKeyFromBase58(string(addr))
}
