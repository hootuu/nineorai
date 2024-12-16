package domains

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/keys"
	"regexp"
	"strings"
)

type StakeAddr = keys.Address

type Stake struct {
	Link      Link         `json:"link" bson:"link"`
	Authority keys.Address `json:"authority" bson:"authority"`
	Network   NetworkAddr  `json:"network" bson:"network"`
	Symbol    StakeSymbol  `json:"symbol" bson:"symbol"`
	Total     uint64       `json:"total" bson:"total"`
	Ctrl      Ctrl         `json:"ctrl" bson:"ctrl"`
	Tag       Tag          `json:"tag" bson:"tag"`
	Meta      Meta         `json:"meta" bson:"meta"`
}

type StakeAccount struct {
	Mint    StakeAddr `json:"mint" bson:"mint"`
	Balance uint64    `json:"balance" bson:"balance"`
}

type StakeSymbol string

func (s StakeSymbol) Validate() *errors.Error {
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
		return errors.Verify("stake symbol is invalid")
	}
	return nil
}
