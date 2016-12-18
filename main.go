package main

import (
	"fmt"

	invertedindex "github.com/teamelehyean/brill/index"
	repository "github.com/teamelehyean/brill/repository"
	tokenizer "github.com/teamelehyean/brill/tokenizer"
)

func main() {
	repo := repository.NewRepository("C:\\Users\\hackeanwarley\\Go_Works\\src\\github.com\\teamelehyean\\brill\\home")
	var content string
	var index int
	documentNameIDMap := make(map[string]int)
	for repo.HasNext() {
		docName, err := repo.Next()
		documentNameIDMap[docName] = index
		fmt.Println(docName)
		if err != nil {
			// do nothing
		}
		content, err = repo.Get()
		tokensCountMap := tokenizer.GetTokens(content)
		for key, value := range tokensCountMap {
			invertedindex.AddToIndex(key, index, value)
		}
		index++
		fmt.Printf("\n\n")
	}
	invertedindex.DisplayInvertedIndex()
	// fmt.Println(documentNameIDMap)
}
