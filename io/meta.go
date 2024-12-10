package io

import "nineorai/domains"

type MetaManager interface {
	SetMeta(dict domains.Dict, ctx Ctx) Error
	RemoveMeta(keys []string, ctx Ctx) Error
}
