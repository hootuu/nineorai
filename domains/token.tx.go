package domains

import (
	"github.com/hootuu/nineorai/keys"
	"time"
)

type TokenTx struct {
	Signature    Signature    `json:"signature" bson:"signature"`
	Mint         keys.Address `json:"mint" bson:"mint"`
	Time         time.Time    `json:"time" bson:"time"`
	Action       string       `json:"action" bson:"action"`
	From         keys.Address `json:"from" bson:"from"`
	To           keys.Address `json:"to" bson:"to"`
	Amount       uint64       `json:"amount" bson:"amount"`
	Memo         Memo         `json:"memo" bson:"memo"`
	AfterBalance uint64       `json:"after_balance" bson:"after_balance"`
}
