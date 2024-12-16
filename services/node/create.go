package node

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type Create struct {
	Link      domains.Link        `json:"link" bson:"link"`
	Authority keys.Address        `json:"authority" bson:"authority"`
	Network   domains.NetworkAddr `json:"network" bson:"network"`
	Address   domains.NodeAddr    `json:"address" bson:"address"`
	Ctrl      domains.Ctrl        `json:"ctrl,omitempty" bson:"ctrl,omitempty"`
	Tag       domains.Tag         `json:"tag,omitempty" bson:"tag,omitempty"`
	Meta      domains.Meta        `json:"meta,omitempty" bson:"meta,omitempty"`
}

type CreateResult struct {
	Address domains.NetworkAddr `json:"address" bson:"address"`
}

func (c Create) Validate() *errors.Error {
	if !c.Link.IsValid() {
		return errors.Verify("invalid link")
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
	return nil
}
