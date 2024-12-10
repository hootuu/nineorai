package domains

import "github.com/hootuu/nineorai/keys"

type TokenAddr = keys.Address

type Token struct {
	Authority keys.Address `json:"authority" bson:"authority"`
	Network   NetworkAddr  `json:"network" bson:"network"`
	Symbol    string       `json:"symbol" bson:"symbol"`
	Decimals  uint8        `json:"decimals" bson:"decimals"`
	Ctrl      Ctrl         `json:"ctrl" bson:"ctrl"`
	Tag       Tag          `json:"tag" bson:"tag"`
	Meta      Meta         `json:"meta" bson:"meta"`
}

type TokenAccount struct {
	Mint     TokenAddr `json:"mint" bson:"mint"`
	Amount   string    `json:"amount" bson:"amount"`
	Decimals uint8     `json:"decimals" bson:"decimals"`
	UiAmount string    `json:"ui_amount" bson:"ui_amount"`
}
