package domains

import "github.com/hootuu/nineorai/keys"

type AssetAddr = keys.Address

type Asset struct {
	Authority keys.Address `json:"authority" bson:"authority"`
	Network   NetworkAddr  `json:"network" bson:"network"`
	Symbol    string       `json:"symbol" bson:"symbol"`
	Ctrl      Ctrl         `json:"ctrl" bson:"ctrl"`
	Tag       Tag          `json:"tag" bson:"tag"`
	Meta      Meta         `json:"meta" bson:"meta"`
}

type AssetAccount struct {
	Asset    AssetAddr `json:"asset" bson:"asset"`
	Quantity uint64    `json:"quantity" bson:"quantity"`
}
