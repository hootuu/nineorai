package io

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/keys"
)

type Account struct {
	Address keys.Address `json:"address" bson:"address"`
	Payer   bool         `json:"payer" bson:"payer"`
	Signer  bool         `json:"signer" bson:"signer"`
}

type AccountSet []Account

func NewAccountSet() AccountSet {
	return make(AccountSet, 0)
}

func (set *AccountSet) AddAccount(a Account) *errors.Error {
	existsPayer := false
	for _, acc := range *set {
		if acc.Address == a.Address {
			return nil
		}
		if acc.Payer {
			existsPayer = true
		}
	}
	if existsPayer && a.Payer {
		return errors.Verify("payer account already exists")
	}
	*set = append(*set, a)
	return nil
}

func (set *AccountSet) Add(addr keys.Address) *errors.Error {
	return set.AddAccount(Account{
		Address: addr,
		Payer:   false,
		Signer:  false,
	})
}

func (set *AccountSet) AddPayer(addr keys.Address) *errors.Error {
	return set.AddAccount(Account{
		Address: addr,
		Payer:   true,
		Signer:  true,
	})
}

func (set *AccountSet) AddSigner(addr keys.Address) *errors.Error {
	return set.AddAccount(Account{
		Address: addr,
		Payer:   false,
		Signer:  true,
	})
}

func (set *AccountSet) GetPayer() *Account {
	for _, acc := range *set {
		if acc.Payer {
			return &acc
		}
	}
	return nil
}

func (set *AccountSet) MustGetPayer() (*Account, *errors.Error) {
	payer := set.GetPayer()
	if payer == nil {
		return nil, errors.Verify("payer account not found")
	}
	return payer, nil
}

func (set *AccountSet) HasPayer(addr keys.Address) bool {
	for _, acc := range *set {
		if acc.Address == addr {
			return acc.Payer
		}
	}
	return false
}

func (set *AccountSet) HasSigner(addr keys.Address) bool {
	for _, acc := range *set {
		if acc.Address == addr {
			return acc.Signer || acc.Payer
		}
	}
	return false
}

func (set *AccountSet) ExistsSignerAtLeast() bool {
	c := 0
	for _, acc := range *set {
		if acc.Payer || acc.Signer {
			c++
		}
	}
	return c > 0
}
