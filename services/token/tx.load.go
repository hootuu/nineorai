package token

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/io/pagination"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type TxLoad struct {
	Authority keys.Address     `json:"authority"`
	Mint      keys.Address     `json:"mint"`
	Page      *pagination.Page `json:"page" bson:"page"`
}

func (tx TxLoad) Validate() *errors.Error {
	if tx.Authority.IsEmpty() {
		return errors.Verify("require tx.authority")
	}
	if tx.Mint.IsEmpty() {
		return errors.Verify("require tx.mint")
	}
	if tx.Page == nil {
		tx.Page = pagination.PageALL()
	}
	return nil
}

type TxLoadResult struct {
	Data   []*domains.TokenTx `json:"data" bson:"data"`
	Paging pagination.Paging  `json:"paging" bson:"paging"`
}
