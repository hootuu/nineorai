package participant

import "nineorai/keys"

type Create struct {
}

type CreateCtx struct {
	Payer     keys.Wallet `bson:"payer" json:"payer"`
	Authority keys.Signer `bson:"authority" json:"authority"`
}
