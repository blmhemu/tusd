package main

// This is very simple for now and will be enhanced as needed.

import (
	"os"
	"path"
)

var dataDir string

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dataDir = path.Join(wd, "tus_data")
	if err := os.MkdirAll(dataDir, 0777); err != nil {
		panic(err)
	}
}

func dataPath(fileId string) string {
	return path.Join(dataDir, fileId)
}

func logPath(fileId string) string {
	return dataPath(fileId)+".log"
}

func initDataFile(id string, size int) {
}