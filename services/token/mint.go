package token

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type Mint struct {
	Token   domains.TokenAddr    `json:"token" bson:"token"`
	Receive domains.IdentityAddr `json:"receive" bson:"receive"`
	Amount  uint64               `json:"amount" bson:"amount"`
	Memo    domains.Memo         `json:"memo" bson:"memo"`

	TokenAuthority keys.Address `json:"token_authority" bson:"token_authority"`
}

type MintResult struct {
	Signature domains.Signature `json:"signature" bson:"signature"`
}

func (m Mint) Validate() *errors.Error {
	if m.Token.IsEmpty() {
		return errors.Verify("require token")
	}
	if m.Receive.IsEmpty() {
		return errors.Verify("require receiver")
	}
	if m.Amount == 0 {
		return errors.Verify("require amount")
	}
	if len(m.Memo) == 0 {
		return errors.Verify("require memo")
	}
	return nil
}
