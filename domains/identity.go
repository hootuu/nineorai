package domains

import (
	"nineorai/keys"
)

type IdentityAddr = keys.Address

type Identity struct {
	Symbol string `bson:"symbol" json:"symbol"`
	Ctrl   Ctrl   `bson:"ctrl" json:"ctrl"`
	Tag    Tag    `bson:"tag" json:"tag"`
	Meta   Meta   `bson:"meta" json:"meta"`
}

type IdentityCollar struct {
	Ctrl bool `bson:"ctrl" json:"ctrl"`
	Tag  bool `bson:"tag" json:"tag"`
	Meta bool `bson:"meta" json:"meta"`
}

func NewIdentityCollar() *IdentityCollar {
	return &IdentityCollar{}
}

func (collar *IdentityCollar) WithCtrl() *IdentityCollar {
	collar.Ctrl = true
	return collar
}

func (collar *IdentityCollar) WithTag() *IdentityCollar {
	collar.Tag = true
	return collar
}

func (collar *IdentityCollar) WithMeta() *IdentityCollar {
	collar.Meta = true
	return collar
}

type IdentityWrap struct {
	NineoraID keys.NineoraID `bson:"nineora_id" json:"nineora_id"`
	Address   keys.Address   `bson:"address" json:"address"`
	Identity  Identity       `bson:"identity" json:"identity"`
}
