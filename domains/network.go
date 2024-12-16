package domains

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/keys"
	"regexp"
	"strings"
)

type NetworkAddr = keys.Address

type NetworkSymbol string

type Network struct {
	Authority keys.Address `json:"authority" bson:"authority"`
	Symbol    string       `json:"symbol" bson:"symbol"`
	Ctrl      Ctrl         `json:"ctrl" bson:"ctrl"`
	Tag       Tag          `json:"tag" bson:"tag"`
	Meta      Meta         `json:"meta" bson:"meta"`
}

func (s NetworkSymbol) Validate() *errors.Error {
	str := strings.TrimSpace(string(s))
	strLen := len(str)
	if strLen == 0 {
		return errors.Verify("symbol is empty")
	}
	if strLen > 16 {
		return errors.Verify("symbol is too long")
	}

	re := regexp.MustCompile("^[a-zA-Z0-9]{1,16}$")

	b := re.MatchString(str)
	if !b {
		return errors.Verify("network symbol is invalid")
	}
	return nil
}
