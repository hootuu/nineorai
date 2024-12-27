package token

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type Transfer struct {
	FromAddress keys.Address `json:"from_address" bson:"from_address"`
	ToAddress   keys.Address `json:"to_address" bson:"to_address"`
	Authority   keys.Address `json:"authority" bson:"authority"`
	Amount      uint64       `json:"amount" bson:"amount"`
	Memo        domains.Memo `json:"memo" bson:"memo"`
}

type TransferResult struct {
	Signature domains.Signature `json:"signature" bson:"signature"`
}

func (t Transfer) Validate() *errors.Error {
	if t.FromAddress.IsEmpty() {
		return errors.Verify("require from address")
	}
	if t.ToAddress.IsEmpty() {
		return errors.Verify("require to address")
	}
	if t.Authority.IsEmpty() {
		return errors.Verify("require authority")
	}
	if t.FromAddress == t.ToAddress {
		return errors.Verify("from and to address same")
	}
	if t.Amount == 0 {
		return errors.Verify("amount is zero")
	}
	if t.Memo == nil {
		return errors.Verify("require memo")
	}
	if err := t.Memo.MustExists(domains.MemoMemo); err != nil {
		return err
	}
	return nil
}
