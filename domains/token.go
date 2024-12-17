package domains

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/keys"
	"regexp"
	"strings"
)

type TokenAddr = keys.Address

type Token struct {
	Link      Link         `json:"link" bson:"link"`
	Authority keys.Address `json:"authority" bson:"authority"`
	Network   NetworkAddr  `json:"network" bson:"network"`
	Mint      TokenAddr    `json:"address" bson:"address"`
	Symbol    string       `json:"symbol" bson:"symbol"`
	Decimals  uint8        `json:"decimals" bson:"decimals"`
	Ctrl      Ctrl         `json:"ctrl" bson:"ctrl"`
	Tag       Tag          `json:"tag" bson:"tag"`
	Meta      Meta         `json:"meta" bson:"meta"`
}

type TokenAccountAddr = keys.Address

type TokenAccount struct {
	Address  TokenAccountAddr `json:"address" bson:"address"`
	Mint     TokenAddr        `json:"mint" bson:"mint"`
	Symbol   TokenSymbol      `json:"symbol" bson:"symbol"`
	Balance  uint64           `json:"balance" bson:"balance"`
	Decimals uint8            `json:"decimals" bson:"decimals"`
	UiAmount string           `json:"ui_amount" bson:"ui_amount"`
}

type TokenSymbol string

func (s TokenSymbol) Validate() *errors.Error {
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
		return errors.Verify("token symbol is invalid")
	}
	return nil
}
