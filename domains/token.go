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

	TokenCtrlSTOC  = 10 // Short-term order commission
	TokenCtrlMTRPD = 11 // Mid-term reward pool dividend
	TokenCtrlLTEI  = 12 // Long-term equity incentive
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
	Address   TokenAccountAddr `json:"address" bson:"address"`
	Authority keys.Address     `json:"authority" bson:"authority"`
	Mint      TokenAddr        `json:"mint" bson:"mint"`
	Symbol    TokenSymbol      `json:"symbol" bson:"symbol"`
	Supply    uint64           `json:"supply" bson:"supply"`
	Balance   uint64           `json:"balance" bson:"balance"`
	Decimals  uint8            `json:"decimals" bson:"decimals"`
	MintMeta  Meta             `json:"mint_meta" bson:"mint_meta"`
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

const AtaLinkTpl = "_*_*A*T*A*_*_"

func GetAtaLink(mint TokenAddr, authority IdentityAddr) Link {
	return GetTokenAccountLink(AtaLinkTpl, mint, authority)
}

func GetTokenAccountLink(linkStr string, mint TokenAddr, authority keys.Address) Link {
	str := fmt.Sprintf("%s*@*%s*@*%s", linkStr, authority.String(), mint.String())
	return NewLink(str)
}
