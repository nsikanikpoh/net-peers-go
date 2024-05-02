package store

import (
	"log"
	"os"
	"time"
)

type FileInfo struct {
	IsDir   bool
	ModTime time.Time
	Name    string
	Size    int64
}

//Searches for files in stored. 
func GetAllLocalFiles() []FileInfo {
	files, err := os.ReadDir("files/stored")
	if err != nil {
		log.Fatal(err)
	}
	fileNames := make([]FileInfo, 0)
	for _, file := range files {
		fileInfo, err := os.Stat("files/stored/" + file.Name())
		if err == nil{
			if len(file.Name()) >= 64 {
				fileNames = append(fileNames, 
					FileInfo{IsDir: fileInfo.IsDir(), 
						ModTime: fileInfo.ModTime(), 
						Name: fileInfo.Name(), 
						Size: fileInfo.Size()})
			}
		}
	}
	return fileNames
}
