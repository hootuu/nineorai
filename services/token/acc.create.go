package token

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type AccountCreate struct {
	Link      domains.Link      `json:"link" bson:"link"`
	Authority keys.Address      `json:"authority" bson:"authority"`
	Mint      domains.TokenAddr `json:"mint" bson:"mint"`
	Ctrl      domains.Ctrl      `json:"ctrl" bson:"ctrl"`
	Tag       domains.Tag       `json:"tag" bson:"tag"`
	Meta      domains.Meta      `json:"meta" bson:"meta"`
}

func (ac AccountCreate) Validate() *errors.Error {
	if !ac.Link.IsValid() {
		return errors.Verify("invalid token account link")
	}
	if ac.Authority.IsEmpty() {
		return errors.Verify("invalid token account authority")
	}
	if ac.Mint.IsEmpty() {
		return errors.Verify("invalid token mint")
	}
	return nil
}

type AccountCreateResult struct {
	Address keys.Address `json:"address" bson:"address"`
}
