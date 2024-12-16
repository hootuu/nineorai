package main

import (
	"fmt"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
)

type User struct {
	Name  string
	Age   int
	Admin bool
}

func (u User) Validate() *errors.Error {
	return nil
}

func main() {
	req := io.NewRequest(&User{
		Name:  "demo",
		Age:   20,
		Admin: true,
	})
	aKey, _ := keys.NewKey()
	bKey, _ := keys.NewKey()
	cKey, _ := keys.NewKey()
	_ = req.Accounts.AddAccount(io.Account{
		Address: aKey.Public.Address(),
		Payer:   false,
		Signer:  false,
	})
	_ = req.Accounts.AddAccount(io.Account{
		Address: bKey.Public.Address(),
		Payer:   true,
		Signer:  false,
	})
	_ = req.Accounts.AddAccount(io.Account{
		Address: cKey.Public.Address(),
		Payer:   false,
		Signer:  true,
	})
	err := req.Sign(map[keys.Address]keys.PrivateKey{
		bKey.Public.Address(): bKey.Private,
		cKey.Public.Address(): cKey.Private,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonBytes, err := req.Marshal()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonBytes))

	revReq, err := io.UnmarshalRequest[User](jsonBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(revReq)
}
