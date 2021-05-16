package src

import (
	"fmt"
	"os"
)

//It will store all the results in analytics/result folder
type Filesystem struct {
}

func (fs Filesystem) OpenFile(fileName string) *os.File {
	file, err := os.OpenFile("result/"+fileName+".txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to open success.txt", " err: ", err.Error())
	}
	return file
}

func (fs Filesystem) CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Unable to close file: "+file.Name(), " err: ", err.Error())
	}
}

func (fs Filesystem) Write(file *os.File, s string) {
	_, err := file.WriteString(s + "\n")
	if err != nil {
		fmt.Println("Failed to write to success.txt", " err: ", err.Error())
	}
}

func NewFilesystem() Filesystem {
	return Filesystem{}
}
