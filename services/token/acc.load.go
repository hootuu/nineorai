package token

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type AccLoadByAuth struct {
	Authority domains.IdentityAddr  `json:"authority" bson:"authority"`
	Networks  []domains.NetworkAddr `json:"network" bson:"network"` //IF EMPTY, RETURN ALL
	Page      pagination.Page       `json:"page" bson:"page"`
}

func (l AccLoadByAuth) Validate() *errors.Error {
	if l.Authority.IsEmpty() {
		return errors.Verify("invalid authority")
	}
	if l.Page.Size <= 0 {
		l.Page.Size = pagination.DefaultPagingSize
	}
	if l.Page.Numb <= 0 {
		l.Page.Numb = pagination.FirstPage
	}
	return nil
}

type AccLoadResult struct {
	Data   []*domains.TokenAccount `json:"data" bson:"data"`
	Paging pagination.Paging       `json:"paging" bson:"paging"`
}

func (r *AccLoadResult) One() *domains.TokenAccount {
	if r.Data == nil || len(r.Data) == 0 {
		return nil
	}
	return r.Data[0]
}

type AccLoadByLink struct {
	Link      string       `json:"link" bson:"link"`
	Mint      keys.Address `json:"mint" bson:"mint"`
	Authority keys.Address `json:"authority" bson:"authority"`
}

func (l AccLoadByLink) Validate() *errors.Error {
	if l.Link == "" {
		return errors.Verify("invalid link")
	}
	if l.Mint.IsEmpty() {
		return errors.Verify("invalid mint")
	}
	if l.Authority.IsEmpty() {
		return errors.Verify("invalid authority")
	}
	return nil
}
