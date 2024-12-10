package asset

import (
	"nineorai/domains"
	"nineorai/io"
)

type Service interface {
	io.CtrlManager
	io.TagManager
	io.MetaManager

	Create(req Create, ctx CreateCtx) (domains.AssetAddr, io.Error)
}
