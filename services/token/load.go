package token

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/hootuu/nineorai/domains"
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
