package helper

import "database/sql"

// CommitOrRollback is a helper function to commit or rollback a transaction based on whether a panic occurred.
func CommitOrRollback(tx *sql.Tx) {
	defer func() {
		err := recover()
		if err != nil {
			errorRollback := tx.Rollback()
			PanicIfError(errorRollback)
			panic(err)
		} else {
			errorCommit := tx.Commit()
			PanicIfError(errorCommit)
		}
	}()
}
