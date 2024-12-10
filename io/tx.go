package io

type Tx struct {
	Tx    string `bson:"tx" json:"tx"`
	Error Error  `bson:"error" json:"error"`
}
