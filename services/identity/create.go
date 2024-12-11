package identity

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/strs"
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
)

type Create struct {
	CustomID   string               `bson:"custom_id" json:"custom_id"`
	Password   []byte               `bson:"password" json:"password"`
	Trustee    bool                 `bson:"trustee" json:"trustee"`
	Address    domains.IdentityAddr `bson:"address" json:"address"`
	PrivateKey []byte               `bson:"private_key" json:"private_key"`
}

type CreateCtx struct {
	Payer keys.PrivateKey `bson:"payer" json:"payer"`
}

func (req *Create) Validate() io.Error {
	if strs.IsEmpty(req.CustomID) {
		return errors.E("require_custom_id", "custom_id is required")
	}
	if len(req.Password) == 0 {
		return errors.E("require_password", "password is required")
	}
	if !req.Trustee {
		if len(req.PrivateKey) == 0 {
			return errors.E("require_private_key", "private_key is required when trustee is false")
		}
	} else {
		if len(req.Address) == 0 {
			return errors.E("require_address", "address is required when trustee")
		}
	}
	return nil
}
