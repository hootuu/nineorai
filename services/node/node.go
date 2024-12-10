package node

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
)

type Service interface {
	Create(node domains.ValuableNode, authority keys.PrivateKey) (keys.Address, io.Error)
	SetSuperior(addr keys.Address, superior keys.Address, auth keys.PrivateKey) io.Error
	io.CtrlManager
	io.TagManager
	io.MetaManager
}
