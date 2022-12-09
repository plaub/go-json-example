package main

import "testing"

func TestReadFilesFromPath(t *testing.T) {
	path := "./testdata/"
	files, err := readFilesFromPath(path)

	if err != nil {
		t.Error(err)
	}

	if len(files) == 0 {
		t.Error("No files found.")
	}
}

func TestReadSportsDataFromFiles(t *testing.T) {
	path := "./testdata/"
	files, err := readFilesFromPath(path)

	if err != nil {
		t.Error(err)
	}

	sportsDataArr, err := readSportsDataFromFiles(path, files)

	if err != nil {
		t.Error(err)
	}

	if len(sportsDataArr) == 0 {
		t.Error("No sports data found.")
	}
}
