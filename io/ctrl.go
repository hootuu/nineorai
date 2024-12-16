package io

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
)

type CtrlManager interface {
	SetCtrl(ctrl domains.Ctrl, ctx Ctx) *errors.Error
}
