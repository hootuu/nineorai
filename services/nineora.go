package services

import (
	"github.com/hootuu/nineorai/services/asset"
	"github.com/hootuu/nineorai/services/identity"
	"github.com/hootuu/nineorai/services/network"
	"github.com/hootuu/nineorai/services/node"
	"github.com/hootuu/nineorai/services/stake"
	"github.com/hootuu/nineorai/services/token"
	"github.com/hootuu/nineorai/services/trustee"
)

type Nineora interface {
	Trustee() trustee.Service
	Identity() identity.Service
	Network() network.Service
	Node() node.Service
	Token() token.Service
	Stake() stake.Service
	Asset() asset.Service
}
