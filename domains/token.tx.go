package domains

import (
	"github.com/hootuu/nineorai/keys"
	"time"
)

type TokenTx struct {
	Signature Signature    `json:"signature" bson:"signature"`
	Mint      keys.Address `json:"mint" bson:"mint"`
	Time      time.Time    `json:"time" bson:"time"`
	Action    string       `json:"action" bson:"action"`
	From      keys.Address `json:"from" bson:"from"`
	FromAfter uint64       `json:"from_after" bson:"from_after"`
	To        keys.Address `json:"to" bson:"to"`
	ToAfter   uint64       `json:"to_after" bson:"to_after"`
	Amount    uint64       `json:"amount" bson:"amount"`
	Memo      Memo         `json:"memo" bson:"memo"`
}
