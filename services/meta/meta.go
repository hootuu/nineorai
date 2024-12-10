package meta

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	nio "io"
)

type Service interface {
	Create(meta domains.Meta) (domains.MetaID, domains.MetaRef, io.Error)
	Save(data nio.Reader) (domains.MetaID, domains.MetaRef, io.Error)
}
