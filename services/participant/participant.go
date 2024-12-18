package participant

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
)

type Service interface {
	io.CtrlManager
	io.TagManager
	io.MetaManager

	Create(req Create, ctx CreateCtx) (domains.ParticipantAddr, io.Error)
}
