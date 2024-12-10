package io

import "nineorai/domains"

type CtrlManager interface {
	SetCtrl(ctrl domains.Ctrl, ctx Ctx) Error
}
