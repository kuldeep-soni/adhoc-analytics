package src

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Analysis1(t *testing.T) {
	dbManager := NewDBManager()
	dbManager.Initialise()
	filesystem := NewFilesystem()

	var wg sync.WaitGroup
	go Execute(&wg, Process, dbManager, filesystem)
	wg.Wait()
	fmt.Println("Complete")
}

func Execute(wg *sync.WaitGroup, process func(manager DBManager, filesystem Filesystem), manager DBManager, filesystem Filesystem) {
	wg.Add(1)
	defer wg.Done()
	process(manager, filesystem)
}

func Process(dbManager DBManager, fs Filesystem) {
	//dbManager.ClientDatabase.Select()
	//dbManager.BillingDatabase.Select()
	//dbManager.CacheDatabase.HGet()

	resultObject := Result1{}
	resultObject.ToJson()
	resultObject.ToDollarSV()

	file := fs.OpenFile("location")
	fs.Write(file, resultObject.ToJson())
	fs.CloseFile(file)
}
