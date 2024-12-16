package io

import (
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/nineorai/domains"
)

type TagManager interface {
	AddTag(tag domains.Tag, ctx Ctx) *errors.Error
	RemoveTag(tag domains.Tag, ctx Ctx) *errors.Error
}
