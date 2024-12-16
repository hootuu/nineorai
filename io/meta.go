package io

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
)

type MetaManager interface {
	SetMeta(dict domains.Dict, ctx Ctx) *errors.Error
	RemoveMeta(keys []string, ctx Ctx) *errors.Error
}
