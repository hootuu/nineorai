package io

import "nineorai/domains"

type TagManager interface {
	AddTag(tag domains.Tag, ctx Ctx) Error
	RemoveTag(tag domains.Tag, ctx Ctx) Error
}
