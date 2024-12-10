package domains

import "nineorai/keys"

type ParticipantAddr = keys.Address

type Participant struct {
	Network keys.Address `json:"network" bson:"network"`
	Node    keys.Address `json:"node" bson:"node"`
	Ctrl    Ctrl         `json:"ctrl" bson:"ctrl"`
	Tag     Tag          `json:"tag" bson:"tag"`
	Meta    Meta         `json:"meta" bson:"meta"`
}
