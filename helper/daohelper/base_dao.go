package daohelper

import "xcontrib/rename/ent"

type BaseDao struct {
	client func() *ent.Client
	hooks  []ent.Hook
}

func (d BaseDao) Client() *ent.Client {
	return EntClient()
}

func (d *BaseDao) SetHooks(hook []ent.Hook) {
	d.hooks = hook
}

func (d BaseDao) Hooks() []ent.Hook {
	return d.hooks
}
