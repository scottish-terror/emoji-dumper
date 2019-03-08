package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type emoji struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {

	var tmpEmoji emoji
	var masterEmoji []emoji
	var files []string

	path := flag.String("p", "", "Path to read")
	flag.Parse()

	if *path == "" {
		fmt.Println("-p required!")
	}

	myPath := *path

	root := myPath
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "" {
			fmt.Println("Skipping invalid file with no extension: " + info.Name())
		} else {
			files = append(files, info.Name())
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		tmpEmoji.URL = myPath + "\\" + file
		tmpEmoji.Name = file[0 : len(file)-4]
		masterEmoji = append(masterEmoji, tmpEmoji)
	}

	writeIt, _ := json.MarshalIndent(masterEmoji, "", " ")
	_ = ioutil.WriteFile("output.json", writeIt, 0644)
}
