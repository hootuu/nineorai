package node

import (
	"nineorai/domains"
	"nineorai/io"
	"nineorai/keys"
)

type Service interface {
	Create(node domains.ValuableNode, authority keys.PrivateKey) (keys.Address, io.Error)
	SetSuperior(addr keys.Address, superior keys.Address, auth keys.PrivateKey) io.Error
	io.CtrlManager
	io.TagManager
	io.MetaManager
}
