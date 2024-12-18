package domains

import (
	"fmt"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/keys"
	"math"
	"regexp"
	"strings"
)

const (
	TokenCtrlNonDivisible    = 1
	TokenCtrlNonTransferable = 2
	TokenCtrlBurnable        = 3
	TokenCtrlNonCollateral   = 4
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
	Supply   uint64           `json:"supply" bson:"supply"`
	Balance  uint64           `json:"balance" bson:"balance"`
	Decimals uint8            `json:"decimals" bson:"decimals"`
}

func (t *TokenAccount) UiAmount() string {
	if t.Decimals == 0 {
		return fmt.Sprintf("%d", t.Balance)
	}
	divisor := math.Pow(10, float64(t.Decimals))
	amount := float64(t.Balance) / divisor
	return fmt.Sprintf(fmt.Sprintf("%%.%df", t.Decimals), amount)
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
