package node

import (
	"nineorai/domains"
	"nineorai/keys"
)

type Create struct {
	Payer     keys.PublicKey       `bson:"payer" json:"payer"`
	Authority keys.PrivateKey      `bson:"authority" json:"authority"`
	Node      domains.ValuableNode `bson:"node" json:"node"`
}
