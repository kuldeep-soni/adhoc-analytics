package analytics

import "adhoc-analytics/src"

//call this function before executing the process in test file. See query1_test.go for example
func initialise() (dbManager src.DBManager, filesystem src.Filesystem) {
	dbManager = src.NewDBManager()
	dbManager.Initialise()
	filesystem = src.NewFilesystem()
	return
}
