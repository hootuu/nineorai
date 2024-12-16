package network

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type Create struct {
	Link      domains.Link          `json:"link" bson:"link"`
	Authority keys.Address          `json:"authority" bson:"authority"`
	Address   domains.NetworkAddr   `json:"address" bson:"address"`
	Symbol    domains.NetworkSymbol `json:"symbol" bson:"symbol"`
	Ctrl      domains.Ctrl          `json:"ctrl,omitempty" bson:"ctrl,omitempty"`
	Tag       domains.Tag           `json:"tag,omitempty" bson:"tag,omitempty"`
	Meta      domains.Meta          `json:"meta,omitempty" bson:"meta,omitempty"`
}

type CreateResult struct {
	Address domains.NetworkAddr `json:"address"`
}

func (c Create) Validate() *errors.Error {
	if c.Authority.IsEmpty() {
		return errors.Verify("require authority")
	}
	if c.Address.IsEmpty() {
		return errors.Verify("require address")
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
	return nil
}
