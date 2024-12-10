package token

import "nineorai/domains"

type Create struct {
	Token domains.Token `bson:"token" json:"token"`
}
