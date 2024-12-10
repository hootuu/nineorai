package domains

import "nineorai/keys"

type StakeAddr = keys.Address

type Stake struct {
	Authority keys.Address `json:"authority" bson:"authority"`
	Network   NetworkAddr  `json:"network" bson:"network"`
	Symbol    string       `json:"symbol" bson:"symbol"`
	Total     uint64       `json:"total" bson:"total"`
	Ctrl      Ctrl         `json:"ctrl" bson:"ctrl"`
	Tag       Tag          `json:"tag" bson:"tag"`
	Meta      Meta         `json:"meta" bson:"meta"`
}

type StakeAccount struct {
	Mint    StakeAddr `json:"mint" bson:"mint"`
	Balance uint64    `json:"balance" bson:"balance"`
}
