package main

import (
	"Gemini-D/fetch-github-stars/format"
	"Gemini-D/fetch-github-stars/repo"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var list []*repo.Repo

	for i, v := range os.Args {
		if i > 0 {
			response, err := http.Get("https://api.github.com/repos/" + v)
			if err != nil {
				continue
			}

			body, _ := io.ReadAll(response.Body)
			repo := new(repo.Repo)
			err = json.Unmarshal(body, repo)
			if err != nil {
				fmt.Print(err)
				continue
			}

			list = append(list, repo)
		}
	}

	format.PrintTable(list)
}
