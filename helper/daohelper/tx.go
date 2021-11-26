package daohelper

import (
	"context"
	"xcontrib/rename/ent"
	"github.com/pkg/errors"
)

func WithTx(ctx context.Context, client *ent.Client, fns ...func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	for _, fn := range fns {
		if err := fn(tx); err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
			}
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}
