package vn

import (
	"nineorai/io"
	"nineorai/keys"
)

type Service interface {
	io.CtrlManager
	io.TagManager
	io.MetaManager

	Create(req Create) (keys.Address, io.Error)
}
