package src

import (
	"adhoc-analytics/src/relationalDB"
	"adhoc-analytics/src/relationalDB/postgres"
	_ "github.com/lib/pq"
	"sync"
)

var (
	syncOnce  sync.Once
	dbManager DBManager
)

//You can add mutiple databases here, be it sql/nosql/inmemory db
type DBManager struct {
	ClientDatabase relationalDB.IDatabase
}

//initialise your databases here
func (d *DBManager) Initialise() {
	d.ClientDatabase = postgres.GetPostgresClient(postgres.PostgresConfig{
		DbName:          "your_db_name",
		UserName:        "user_name",
		Host:            "db_host",
		Password:        "password",
		AllowEncryption: false,
		SSLMode:         "disable",
		SSLRootCert:     "root.crt",
		MaxIdleConns:    10,
		MaxOpenConns:    20,
		ConnMaxLifetime: 2000,
	})
}

func initialise() {
	dbManager = DBManager{}
}

func NewDBManager() DBManager {
	syncOnce.Do(initialise)
	return dbManager
}
