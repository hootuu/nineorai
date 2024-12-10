package vn

import (
	"nineorai/domains"
	"nineorai/keys"
)

type Create struct {
	Payer   keys.Wallet     `bson:"payer" json:"payer"`
	Network domains.Network `bson:"network" json:"network"`
}
