package vn

import (
	"github.com/hootuu/nineorai/domains"
	"github.com/hootuu/nineorai/io"
	"github.com/hootuu/nineorai/keys"
)

type Service interface {
	Create(req Create) (keys.Address, io.Error)
	SetCtrl(ctrl domains.Ctrl, ctx io.Ctx) io.Error
	AddTag(tag domains.Tag, ctx io.Ctx) io.Error
	RemoveTag(tag domains.Tag, ctx io.Ctx) io.Error
}
