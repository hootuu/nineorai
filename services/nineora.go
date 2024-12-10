package services

import (
	"nineorai/services/asset"
	"nineorai/services/identity"
	"nineorai/services/node"
	"nineorai/services/stake"
	"nineorai/services/token"
	"nineorai/services/vn"
)

type Nineora interface {
	Identity() identity.Service
	Network() vn.Service
	Node() node.Service
	Token() token.Service
	Stake() stake.Service
	Asset() asset.Service
}
