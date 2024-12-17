package domains

import (
	"github.com/hootuu/nineorai/keys"
)

type IdentityAddr = keys.Address

type Identity struct {
	Address   IdentityAddr   `bson:"address" json:"address"`
	NineoraID keys.NineoraID `bson:"nineora_id" json:"nineora_id"`
	Link      Link           `bson:"link" json:"link"`
	Ctrl      Ctrl           `bson:"ctrl,omitempty" json:"ctrl,omitempty"`
	Tag       Tag            `bson:"tag,omitempty" json:"tag,omitempty"`
	Meta      Meta           `bson:"meta,omitempty" json:"meta,omitempty"`
}
