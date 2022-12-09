package main

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/plaub/go-json-example/sports"
)

func main() {
	if os.Args[1] == "help" {
		fmt.Println("Usage: go run main.go [help|path] [path to data directory]")
		return
	}

	var path string
	if os.Args[1] == "path" {
		path = os.Args[2]
	} else {
		panic("Please provide a path to the data directory.")
	}

	if (path[len(path)-1:]) != "/" {
		path = path + "/"
	}

	files, filesError := readFilesFromPath(path)

	if filesError != nil {
		panic(filesError.Error())
	}

	sportsDataArr, sportsDataError := readSportsDataFromFiles(path, files)

	if sportsDataError != nil {
		panic(sportsDataError.Error())
	}

	printSportsData(sportsDataArr, files)
}

func printSportsData(sportsDataArr []sports.SportsData, files []fs.FileInfo) {
	fmt.Println("Yeah. It worked. :)")
	fmt.Println()

	for index, item := range sportsDataArr {
		fmt.Println("Data points in " + files[index].Name() + ": " + fmt.Sprint(item.GetEntriesLength()))
	}

	fmt.Println()
	fmt.Println("Data Table:")

	sports.SportsDataTablePrint(sportsDataArr)
}

func readSportsDataFromFiles(path string, files []fs.FileInfo) ([]sports.SportsData, error) {
	var sportsDataArr = []sports.SportsData{}

	for _, file := range files {
		sportsData, sportsError := sports.ReadSportsData(path + file.Name())

		if sportsError != nil {
			return nil, errors.New("Could not read " + "'" + file.Name() + "', Error: " + sportsError.Error())
		}

		sportsDataArr = append(sportsDataArr, sportsData)
	}
	return sportsDataArr, nil
}

func readFilesFromPath(path string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return files, nil
}
