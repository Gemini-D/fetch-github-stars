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
	var repos = []string{
		"hyperf/hyperf",
		"easy-swoole/easyswoole",
		"swoft-cloud/swoft",
		"laravel/laravel",
		"Yurunsoft/imi",
		"hyperf/nano",
		"hyperf/hyperf-docker",
		"hyperf/hyperf-skeleton",
	}

	if len(os.Args) > 1 {
		repos = os.Args[1:]
	}

	for _, v := range repos {
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

	format.PrintTable(list)
}
