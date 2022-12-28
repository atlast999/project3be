package transaction

import (
	"context"
	"database/sql"

	db "github.com/atlast999/project3be/db/gen"
)

type TxInstance struct {
	Queries    *db.Queries
	DBInstance *sql.DB
}

func NewTxInstance(dbInstance *sql.DB) *TxInstance {
	return &TxInstance{
		Queries:    db.New(dbInstance),
		DBInstance: dbInstance,
	}
}

func (txInstance *TxInstance) ExecTransaction(ctx context.Context, txBlock func(*db.Queries) error) error {
	tx, err := txInstance.DBInstance.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = txBlock(txInstance.Queries)
	if err != nil {
		return err
	}
	return tx.Commit()
}
