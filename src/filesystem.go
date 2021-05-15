package src

import (
	"os"
)

type Filesystem struct {
}

func (fs Filesystem) OpenFile(absFilePath string) *os.File {
	file, err := os.OpenFile(absFilePath+".txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		//plog.Info("Failed to open success.txt", plog.Params{"err": err})
	}
	return file
}

func (fs Filesystem) CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		//plog.Info("Unable to close file: "+file.Name(), plog.Params{"err": err})
	}
}

func (fs Filesystem) Write(file *os.File, s string) {
	_, err := file.WriteString(s + "\n")
	if err != nil {
		//plog.Info("Failed to write to success.txt", plog.Params{"err": err})
	}
}

func NewFilesystem() Filesystem {
	return Filesystem{}
}
