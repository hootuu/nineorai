package token

import "github.com/hootuu/nineorai/domains"

type Create struct {
	Token domains.Token `bson:"token" json:"token"`
}
