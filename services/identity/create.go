package identity

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
)

type Create struct {
	Payer    keys.PrivateKey  `bson:"payer" json:"payer"`
	Identity domains.Identity `bson:"identity" json:"identity"`
}

func (req *Create) Validate() io.Error {
	b, err := keys.ValidatePrivateKey(req.Payer)
	if err != nil {
		return io.NewApiError("invalid_private_key", err.Error())
	}
	if !b {
		return io.NewApiError("invalid_payer_key", "invalid payer")
	}
	return nil
}
