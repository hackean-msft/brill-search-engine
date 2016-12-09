package main

import (
	"fmt"

	repository "github.com/teamelehyean/brill/repository"
)

func main() {
	repo := repository.NewRepository("C:\\Users\\hackeanwarley\\Go_Works\\src\\github.com\\teamelehyean\\brill\\home")
	var content string
	for repo.HasNext() {
		fileName, err := repo.Next()
		if err != nil {
			// do nothing
		}
		content, err = repo.Get()
		fmt.Println("\n" + fileName)
		fmt.Println(content)
	}

	// file, err = repo.Next()
	// if err != nil {
	// 	// do nothing
	// }
	// fmt.Println(file)

	// file, err = repo.Next()
	// if err != nil {
	// 	// do nothing
	// }
	// fmt.Println(file)
}
