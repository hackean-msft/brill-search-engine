package main

import (
	"fmt"

	repository "github.com/teamelehyean/brill/repository"
)

func main() {
	repo := repository.NewRepository("C:\\Users\\hackeanwarley\\Go_Works\\src\\github.com\\teamelehyean\\brill\\home")
	files := repo.GetFiles()
	for _, file := range files {
		fmt.Println(file)
	}
}
