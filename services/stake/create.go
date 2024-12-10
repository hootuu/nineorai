package stake

import (
	"nineorai/domains"
	"nineorai/keys"
)

type Create struct {
	Stake domains.Stake `bson:"stake" json:"stake"`
}

type CreateCtx struct {
	Payer     keys.Wallet `bson:"payer" json:"payer"`
	Authority keys.Signer `bson:"authority" json:"authority"`
}
