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
	Key      *keys.Key
}

func NewCreate(trustee bool, linkStr string, password string, key *keys.Key) *Create {
	return &Create{
		Trustee:  trustee,
		Link:     domains.NewLink(linkStr),
		Password: domains.NewPassword(password),
		Key:      key,
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
	key, _ := keys.NewKey()
	return &Create{
		Trustee:  true,
		Link:     domains.NewLink(idx.New()),
		Password: domains.NewPassword(pwd),
		Key:      key,
	}
}

type CreateResult struct {
	Address keys.Address `json:"address" bson:"address"`
}

func (c Create) Validate() *errors.Error {
	if !c.Link.IsValid() {
		return errors.Verify("invalid link")
	}
	if !c.Password.IsValid() {
		return errors.Verify("invalid password")
	}
	if c.Key == nil {
		return errors.Verify("invalid key")
	}
	return nil
}
