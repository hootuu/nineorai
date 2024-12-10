package domains

import "nineorai/keys"

type NodeAddr = keys.Address

type ValuableNode struct {
	Network   NetworkAddr  `json:"network" bson:"network"`
	Superior  keys.Address `json:"superior" bson:"superior"`
	Authority keys.Address `json:"authority" bson:"authority"`
	Symbol    string       `json:"symbol" bson:"symbol"`
	Ctrl      Ctrl         `json:"ctrl" bson:"ctrl"`
	Tag       Tag          `json:"tag" bson:"tag"`
	Meta      Meta         `json:"meta" bson:"meta"`
}
