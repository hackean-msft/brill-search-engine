package main

import (
	"fmt"

	repository "github.com/teamelehyean/brill/repository"
	tokenizer "github.com/teamelehyean/brill/tokenizer"
)

func main() {
	repo := repository.NewRepository("C:\\Users\\hackeanwarley\\Go_Works\\src\\github.com\\teamelehyean\\brill\\home")
	var content string
	for repo.HasNext() {
		_, err := repo.Next()
		if err != nil {
			// do nothing
		}
		content, err = repo.Get()
		modifiedString := tokenizer.RemoveUnwantedSymbols(content, []string{",", ".", "}", "|", "{", "\r", "\n", "(", ")", ":", "_", "<", ">", "/", "\\"})
		modifiedString = tokenizer.ToLower(modifiedString)
		fmt.Println(modifiedString)
		// tokens := tokenizer.Tokenize(content)
		// for _, token := range tokens {
		// 	removeNonPrintableCharacters(&token)
		// 	fmt.Println(token)
		// }
	}
}
