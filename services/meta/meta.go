package meta

import (
	nio "io"
	"nineorai/domains"
	"nineorai/io"
)

type Service interface {
	Create(meta domains.Meta) (domains.MetaID, domains.MetaRef, io.Error)
	Save(data nio.Reader) (domains.MetaID, domains.MetaRef, io.Error)
}
