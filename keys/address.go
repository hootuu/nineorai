package keys

import (
	"github.com/hootuu/gelato/errors"
	"strings"
)

type Address string

func (addr Address) PublicKey() (PublicKey, *errors.Error) {
	return PublicKeyFromBase58(string(addr))
}

func (addr Address) IsEmpty() bool {
	return len(strings.TrimSpace(string(addr))) == 0
}

func (addr Address) String() string {
	return string(addr)
}
