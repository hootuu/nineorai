package domains

import (
	"github.com/hootuu/gelato/crtpto/md5x"
)

type Link string

func NewLink(seed string) Link {
	return Link(md5x.MD5(seed))
}

func (link Link) String() string {
	return string(link)
}

func (link Link) Equals(str string) bool {
	return link == NewLink(str)
}

func (link Link) IsValid() bool {
	return md5x.IsMD5(link.String())
}
