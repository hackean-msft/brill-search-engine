package repository

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var index = -1
var folders = 0
var mu sync.Mutex

// Repository for all documents that will be indexed
type Repository struct {
	name  string
	files []string
}

// NewRepository ( The Repo for all files and folders )
func (r Repository) NewRepository(path string) {
	files, err := ioutil.ReadDir(path)
	r.files = make([]string, len(files), len(files))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read directory %s, Reason: %v\n", path, err)
	}
	for _, file := range files {
		absoluteFileName := path + string(filepath.Separator) + file.Name()
		r.files = append(r.files, absoluteFileName)
	}
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

// NextFile (Returns the next file)
func (r Repository) NextFile() (string, error) {
	index++
	if index < len(r.files) {
		return r.files[index], nil
	}
	return "", errors.New("No more file to be read")

}
