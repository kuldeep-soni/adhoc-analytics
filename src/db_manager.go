package src

import "sync"

var (
	syncOnce  sync.Once
	dbManager DBManager
)

type DBManager struct {
	ClientDatabase IDatabase
	BillingDatabase IDatabase
}

func (d *DBManager) Initialise() {

}

func initialise() {
	dbManager = DBManager{}
}

func NewDBManager() DBManager {
	syncOnce.Do(initialise)
	return dbManager
}
