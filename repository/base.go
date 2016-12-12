package repository

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var index = -1

// var folders = 0
// var mu sync.Mutex

// Repository for all documents that will be indexed
type Repository struct {
	name  string
	files []string
}

// NewRepository ( The Repo for all files and folders )
func NewRepository(path string) *Repository {
	files, err := ioutil.ReadDir(path)
	pathFiles := make([]string, 0, len(files))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read directory %s, Reason: %v\n", path, err)
	}
	for _, file := range files {
		absoluteFileName := path + string(filepath.Separator) + file.Name()
		pathFiles = append(pathFiles, absoluteFileName)
	}
	return &Repository{name: path, files: pathFiles}
}

// func listDirContents(path string) {
// 	files, err := ioutil.ReadDir(path)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Could not read directory %s, Reason: %v\n", path, err)
// 	}
// 	for _, file := range files {
// 		if file.IsDir() {

// 			mu.Lock()
// 			folders++
// 			mu.Unlock()

// 			dirname := path + string(filepath.Separator) + file.Name()
// 			go listDirContents(dirname)
// 		} else {
// 			mu.Lock() 
// 			append(file)
// 			mu.Unlock()
// 		}
// 	}
// }

// GetFiles (Returns all the file in a particular repository)
func (r Repository) GetFiles() []string {
	return r.files
}

// HasNext checks if there is more file to be returned
func (r Repository) HasNext() bool {
	temp := index + 1
	return temp < len(r.files)
}

// Next Moves the pointer to the next file and returns the file name
func (r Repository) Next() (string, error) {
	index++
	if index < len(r.files) {
		return r.files[index], nil
	}
	return "", errors.New("End of file reached")
}

// Get Reads the Contents of the current file
func (r Repository) Get() (string, error) {
	suffixes := []string{"pdf", "docx", "ppt", "doc", "txt"}
	for _, suffix := range suffixes {
		fileName := r.files[index]
		if strings.HasSuffix(fileName, suffix) {
			result, err := getContents(fileName, suffix)
			if err != nil {
				return "", err
			}
			return result, nil
		}
	}
	return "", errors.New("Invalid File Format")
}

func getContents(filePath string, suffix string) (string, error) {
	if strings.Compare(suffix, "txt") == 0 {
		result, err := ioutil.ReadFile(filePath)
		if err != nil {
			return "", err
		}
		return string(result), nil
	}
	return "", errors.New("Invalid File Format")
}
