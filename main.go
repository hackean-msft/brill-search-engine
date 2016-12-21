package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	invertedindex "github.com/teamelehyean/brill/index"
	"github.com/teamelehyean/brill/ranker"
	repository "github.com/teamelehyean/brill/repository"
	tokenizer "github.com/teamelehyean/brill/tokenizer"
)

func main() {
	invertedindex.DisplayInvertedIndex()
	http.HandleFunc("/", serverRest)
	http.HandleFunc("/notify", notify)
	http.ListenAndServe("localhost:8800", nil)
	// results := ranker.Rank("Lorem Ipsum discrete mathematics script")
	// fmt.Println(results)
}

func getJSONResponse(query string) ([]byte, error) {
	results := ranker.Rank(query)

	url := "http://localhost:22000/documents"

	data, err := json.Marshal(results)
	if err != nil {
		fmt.Println("Could not marshal data")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	var documents []Payload
	json.Unmarshal(body, &documents)

	return json.MarshalIndent(documents, "", " ")
}

func serverRest(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	fmt.Printf("Query: %s", query)
	response, _ := getJSONResponse(query)
	fmt.Fprintf(w, string(response))
}

func notify(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error")
	}
	var p Payload
	json.Unmarshal(body, &p)
	content, err := repository.GetFileContents(p.FileName)
	tokens := tokenizer.GetTokens(content)
	fmt.Println("Finished Reading " + p.FileName)
	fmt.Println("Started Indexing " + p.FileName)
	for _, token := range tokens {
		invertedindex.AddToIndex(token, p.ID)
	}
	fmt.Println("Finished Indexing " + p.FileName)
}

//Payload represents the result
type Payload struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FileName    string `json:"filename"`
	CoverImage  string `json:"cover_img"`
	Uploaded    int64  `json:"uploaded"`
}
