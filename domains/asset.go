package domains

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/keys"
	"regexp"
	"strings"
)

type AssetAddr = keys.Address

type Asset struct {
	Link      Link         `json:"link" bson:"link"`
	Authority keys.Address `json:"authority" bson:"authority"`
	Network   NetworkAddr  `json:"network" bson:"network"`
	Address   AssetAddr    `json:"address" bson:"address"`
	Symbol    string       `json:"symbol" bson:"symbol"`
	Ctrl      Ctrl         `json:"ctrl" bson:"ctrl"`
	Tag       Tag          `json:"tag" bson:"tag"`
	Meta      Meta         `json:"meta" bson:"meta"`
}

type AssetAccount struct {
	Asset    AssetAddr `json:"asset" bson:"asset"`
	Quantity uint64    `json:"quantity" bson:"quantity"`
}

type AssetSymbol string

func (s AssetSymbol) Validate() *errors.Error {
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
