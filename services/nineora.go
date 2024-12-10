package services

import (
	"github.com/hootuu/nineorai/services/asset"
	"github.com/hootuu/nineorai/services/identity"
	"github.com/hootuu/nineorai/services/node"
	"github.com/hootuu/nineorai/services/stake"
	"github.com/hootuu/nineorai/services/token"
	"github.com/hootuu/nineorai/services/vn"
)

type Nineora interface {
	Identity() identity.Service
	Network() vn.Service
	Node() node.Service
	Token() token.Service
	Stake() stake.Service
	Asset() asset.Service
}
