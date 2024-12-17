package identity

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type Get struct {
	Address domains.IdentityAddr `json:"address" bson:"address"`
}

func (g Get) Validate() *errors.Error {
	if g.Address.IsEmpty() {
		return errors.Verify("require address")
	}
	return nil
}

type GetByLink struct {
	Link domains.Link `json:"link" bson:"link"`
}

func (g GetByLink) Validate() *errors.Error {
	if !g.Link.IsValid() {
		return errors.Verify("invalid link")
	}
	return nil
}

type GetByNineoraID struct {
	NineoraID keys.NineoraID `json:"nineora_id" bson:"nineora_id"`
}

func (g GetByNineoraID) Validate() *errors.Error {
	if !g.NineoraID.IsValid() {
		return errors.Verify("invalid nineora_id")
	}
	return nil
}

type GetResult struct {
	Identity domains.Identity `json:"identity" bson:"identity"`
}
