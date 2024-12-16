package io

import "github.com/hootuu/gelato/errors"

type Tx struct {
	Tx    string        `bson:"tx" json:"tx"`
	Error *errors.Error `bson:"error" json:"error"`
}
