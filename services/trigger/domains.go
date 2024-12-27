package trigger

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
	"regexp"
)

type Trigger struct {
	Contract keys.Address          `json:"contract" bson:"contract"`
	Code     string                `json:"code" bson:"code"`
	Ctx      domains.Dict          `json:"ctx" bson:"ctx"`
	Accounts map[string]io.Account `json:"accounts" bson:"accounts"`
	Memo     domains.Memo          `json:"memo" bson:"memo"`
}

func (t Trigger) Validate() *errors.Error {
	if t.Contract.IsEmpty() {
		return errors.Verify("require contract")
	}
	regex := regexp.MustCompile(`^[A-Z0-9_]{1,32}$`)
	if !regex.MatchString(t.Code) {
		return errors.Verify("invalid trigger.code")
	}
	if len(t.Memo) == 0 {
		return errors.Verify("require memo")
	}
	if err := t.Memo.MustExists(domains.MemoMemo); err != nil {
		return err
	}
	if len(t.Ctx) == 0 {
		return errors.Verify("require trigger.ctx")
	}
	return nil
}

func (t Trigger) AddAccount(key string, acc io.Account) {
	if t.Accounts == nil {
		t.Accounts = make(map[string]io.Account)
	}
	t.Accounts[key] = acc
}

type Result struct {
	Signature domains.Signature `json:"signature" bson:"signature"`
}
