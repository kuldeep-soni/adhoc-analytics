package analytics

import (
	"adhoc-analytics/convertor"
	"adhoc-analytics/src"
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_Analysis1(t *testing.T) {
	dbManager, filesystem := initialise()

	ExecuteProcess(context.Background(), dbManager, filesystem)
	fmt.Println("Complete")
}

func ExecuteProcess(ctx context.Context, manager src.DBManager, filesystem src.Filesystem) {
	//you can write the logic to run the process in multiple go routines. Pass unique work to each go routine.
	//You can have your own parameters to the process
	workIds := []string{"Test1", "Test2"}

	var wg sync.WaitGroup
	wg.Add(2)
	go Process(ctx, &wg, manager, filesystem, workIds[:len(workIds)/2])
	go Process(ctx, &wg, manager, filesystem, workIds[len(workIds)/2:])
	wg.Wait()
}

func Process(ctx context.Context, wg *sync.WaitGroup, dbManager src.DBManager, fs src.Filesystem, workIds []string) {
	defer wg.Done()
	metaData := ResultObjectExample{}

	otherComments := fs.OpenFile("other_logs")
	successFile := fs.OpenFile("success")
	failureFile := fs.OpenFile("failure")

	for _, id := range workIds {
		var exampleDAO DBExampleStruct
		err := dbManager.ClientDatabase.SelectOne(&exampleDAO, "select * from db_example_table where id = $1", id)
		if err != nil {
			metaData.Set(time.Now(), exampleDAO.Id, len(exampleDAO.EntityType), err.Error(), "Failed to fetch data from db")
			fs.Write(failureFile, convertor.ToCSV(metaData))
			continue
		}
		metaData.Set(time.Now(), exampleDAO.Id, len(exampleDAO.EntityType), "", "Success")
		fs.Write(successFile, convertor.ToCSV(metaData))
	}

	metaData.Set(time.Now(), "", -1, "None", "Completed")
	fs.Write(otherComments, convertor.ToJson(metaData))

	fs.CloseFile(successFile)
	fs.CloseFile(failureFile)
	fs.CloseFile(otherComments)
}
