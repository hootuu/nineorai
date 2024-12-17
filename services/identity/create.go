package identity

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type Create struct {
	Link     domains.Link         `json:"link"`
	Password domains.Password     `json:"password"`
	Address  domains.IdentityAddr `json:"address"`
	Ctrl     domains.Ctrl         `json:"ctrl"`
	Tag      domains.Tag          `json:"tag"`
	Meta     domains.Meta         `json:"meta"`
}

type CreateResult struct {
	NineoraID keys.NineoraID       `json:"nineora_id"`
	Address   domains.IdentityAddr `json:"address"`
}

func (c Create) Validate() *errors.Error {
	if !c.Link.IsValid() {
		return errors.Verify("link is required")
	}
	if !c.Password.IsValid() {
		return errors.Verify("password invalid")
	}
	if c.Address.IsEmpty() {
		return errors.Verify("address is required")
	}
	return nil
}
