package helper

import (
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		tx.Rollback()
		//PanicIfError(errorRollback)
		panic(err)
	} else {
		tx.Commit()
		//PanicIfError(errorCommit)
	}
}
