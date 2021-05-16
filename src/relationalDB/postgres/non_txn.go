package postgres

import (
	"adhoc-analytics/src/relationalDB"
	"database/sql"
	"errors"
	"gopkg.in/gorp.v2"
	"time"
)

func GetPostgresClient(config PostgresConfig) relationalDB.IDatabase {
	return &postgresDb{initPostgresDataBase(config)}
}

type postgresDb struct {
	gorpDbMap *gorp.DbMap
}

type PostgresConfig struct {
	DbName          string
	UserName        string
	Host            string
	Password        string
	AllowEncryption bool
	SSLMode         string
	SSLRootCert     string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

func (pg *postgresDb) Insert(list ...interface{}) (err error) {
	return pg.gorpDbMap.Insert(list...)
}

func (pg *postgresDb) Update(list ...interface{}) (i int64, err error) {
	updatedRowsCount, err := pg.gorpDbMap.Update(list...)
	if err != nil {
		return
	}
	if int(updatedRowsCount) < len(list) {
		err = errors.New("Not all rows updated")
	}
	return updatedRowsCount, err
}

func (pg *postgresDb) Delete(list ...interface{}) (i int64, err error) {
	return pg.gorpDbMap.Delete(list...)
}

func (pg *postgresDb) Exec(query string, args ...interface{}) (r sql.Result, err error) {
	return pg.gorpDbMap.Exec(query, args...)
}

func (pg *postgresDb) Select(i interface{}, query string, args ...interface{}) (ri []interface{}, err error) {
	return pg.gorpDbMap.Select(i, query, args...)
}

func (pg *postgresDb) SelectOne(holder interface{}, query string, args ...interface{}) (err error) {
	return pg.gorpDbMap.SelectOne(holder, query, args...)
}

func (pg *postgresDb) Commit() (err error) {
	panic("Commit not supported as this is not txn client")
	return
}

func (pg *postgresDb) Rollback() (err error) {
	panic("Rollback not supported as this is not txn client")
	return
}

func initPostgresDataBase(config PostgresConfig) *gorp.DbMap {
	dbName := config.DbName
	username := config.UserName
	host := config.Host
	password := config.Password
	allowEncryption := config.AllowEncryption
	sslMode := config.SSLMode
	sslRootCert := config.SSLRootCert

	var encryptionString string
	if allowEncryption {
		encryptionString = ` sslrootcert=` + sslRootCert
	}
	db, err := sql.Open("postgres", `dbname=`+dbName+` user=`+username+` password=`+password+` host=`+host+` sslmode=`+sslMode+encryptionString)
	if err != nil {
		panic("Main db connection failed")
	}

	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime))

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	return dbMap
}
