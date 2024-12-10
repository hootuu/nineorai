package vn

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/keys"
)

type Create struct {
	Payer   keys.Wallet     `bson:"payer" json:"payer"`
	Network domains.Network `bson:"network" json:"network"`
}
