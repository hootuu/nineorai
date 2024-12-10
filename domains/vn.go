package domains

import "github.com/hootuu/nineorai/keys"

type NetworkAddr = keys.Address

type Network struct {
	Authority keys.Address `json:"authority" bson:"authority"`
	Symbol    string       `json:"symbol" bson:"symbol"`
	Ctrl      Ctrl         `json:"ctrl" bson:"ctrl"`
	Tag       Tag          `json:"tag" bson:"tag"`
	Meta      Meta         `json:"meta" bson:"meta"`
}
