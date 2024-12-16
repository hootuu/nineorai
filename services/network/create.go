package network

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

const MetaName = "name"
const MetaUri = "uri"

type Create struct {
	Authority keys.Address          `json:"authority"`
	Address   domains.NetworkAddr   `json:"address"`
	Symbol    domains.NetworkSymbol `json:"symbol"`
	Ctrl      domains.Ctrl          `json:"ctrl,omitempty"`
	Tag       domains.Tag           `json:"tag,omitempty"`
	Meta      domains.Meta          `json:"meta,omitempty"`
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
	if err := c.Meta.MustExists(MetaName, MetaUri); err != nil {
		return err
	}
	return nil
}
