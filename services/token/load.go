package token

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/hootuu/nineorai/domains"
)

type LoadByAuthority struct {
	Authority domains.IdentityAddr `json:"authority" bson:"authority"`
	Page      pagination.Page      `json:"page" bson:"page"`
}

func (l LoadByAuthority) Validate() *errors.Error {
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

type LoadResult struct {
	Data   []domains.TokenAccount `json:"data" bson:"data"`
	Paging pagination.Paging      `json:"paging" bson:"paging"`
}
