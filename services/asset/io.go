package asset

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type Create struct {
	Asset domains.Asset `bson:"asset" json:"asset"`
}

type CreateCtx struct {
	Payer     keys.Wallet `bson:"payer" json:"payer"`
	Authority keys.Signer `json:"authority" bson:"authority"`
}
