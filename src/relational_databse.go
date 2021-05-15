package src

import (
	"database/sql"
	"fmt"
)

type IDatabase interface {
	Insert(list ...interface{}) (err error)
	Update(list ...interface{}) (i int64, err error)
	Delete(list ...interface{}) (i int64, err error)
	Exec(query string, args ...interface{}) (r sql.Result, err error)
	Select(i interface{}, query string, args ...interface{}) (ri []interface{}, err error)
	SelectOne(holder interface{}, query string, args ...interface{}) (err error)
	SelectById(holder interface{}, id fmt.Stringer) (err error)
	SelectOneJoin(holder interface{}, query string, args ...interface{}) (err error)
	SelectJoin(holder interface{}, query string, args ...interface{}) (err error)
	Commit() (err error)
	Rollback() (err error)
}
