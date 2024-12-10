package io

import "nineorai/keys"

type Ctx struct {
	Payer     keys.PrivateKey `bson:"payer" json:"payer"`
	Authority keys.PrivateKey `bson:"authority" json:"authority"`
}
