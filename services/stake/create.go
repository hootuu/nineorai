package stake

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type Create struct {
	Link      domains.Link        `json:"link" bson:"link"`
	Authority keys.Address        `json:"authority" bson:"authority"`
	Address   keys.Address        `json:"address" bson:"address"`
	Network   domains.NetworkAddr `json:"network" bson:"network"`
	Symbol    domains.StakeSymbol `json:"symbol" bson:"symbol"`
	Total     uint64              `json:"total" bson:"total"`
	Ctrl      domains.Ctrl        `json:"ctrl" bson:"ctrl"`
	Tag       domains.Tag         `json:"tag" bson:"tag"`
	Meta      domains.Meta        `json:"meta" bson:"meta"`
}

type CreateResult struct {
	Address keys.Address `json:"address" bson:"address"`
}

func (c Create) Validate() *errors.Error {
	if c.Authority.IsEmpty() {
		return errors.Verify("require authority")
	}
	if c.Address.IsEmpty() {
		return errors.Verify("require address")
	}
	if c.Network.IsEmpty() {
		return errors.Verify("require network")
	}
	if err := c.Symbol.Validate(); err != nil {
		return err
	}
	if c.Meta == nil {
		return errors.Verify("require meta")
	}
	if err := c.Meta.MustExists(domains.MetaName, domains.MetaUri); err != nil {
		return err
	}
	if c.Total == 0 {
		return errors.Verify("require total")
	}
	return nil
}
