package trustee

import (
	"github.com/hootuu/gelato/crtpto/rand"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/idx"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type Create struct {
	Trustee  bool
	Link     domains.Link
	Password domains.Password
}

func NewCreate(trustee bool, linkStr string, password string) *Create {
	return &Create{
		Trustee:  trustee,
		Link:     domains.NewLink(linkStr),
		Password: domains.NewPassword(password),
	}
}

func NewRandCreate(trustee bool) *Create {
	if !trustee {
		return &Create{Trustee: false}
	}
	pwd, err := rand.String(16)
	if err != nil {
		pwd = idx.New()
	}
	return &Create{
		Trustee:  true,
		Link:     domains.NewLink(idx.New()),
		Password: domains.NewPassword(pwd),
	}
}

type CreateResult struct {
	Key keys.Key
}

func (c Create) Validate() *errors.Error {
	if !c.Link.IsValid() {
		return errors.Verify("invalid link")
	}
	if !c.Password.IsValid() {
		return errors.Verify("invalid password")
	}
	return nil
}
