package token

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type Create struct {
	Link      domains.Link        `json:"link"`
	Authority keys.Address        `json:"authority"`
	Network   domains.NetworkAddr `json:"network"`
	Address   domains.TokenAddr   `json:"address"`
	Symbol    domains.TokenSymbol `json:"symbol"`
	Decimals  uint8               `json:"decimals"`
	Ctrl      domains.Ctrl        `json:"ctrl"`
	Tag       domains.Tag         `json:"tag"`
	Meta      domains.Meta        `json:"meta"`
}

type CreateResult struct {
	Address domains.TokenAddr `json:"address"`
}

func (c Create) Validate() *errors.Error {
	if !c.Link.IsValid() {
		return errors.Verify("require link")
	}
	if c.Authority.IsEmpty() {
		return errors.Verify("require authority")
	}
	if c.Network.IsEmpty() {
		return errors.Verify("require network")
	}
	if c.Address.IsEmpty() {
		return errors.Verify("require address")
	}
	if err := c.Symbol.Validate(); err != nil {
		return err
	}
	if c.Decimals > 18 || c.Decimals < 0 {
		return errors.Verify("decimals must be between 0 and 18")
	}
	return nil
}
