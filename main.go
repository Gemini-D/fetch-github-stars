package main

import (
	"Gemini-D/fetch-github-stars/format"
	"Gemini-D/fetch-github-stars/repo"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

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
		request, _ := http.NewRequest("GET", "https://api.github.com/repos/"+v, nil)
		request.Header.Set("Authorization", "Token "+token)
		response, err := (&http.Client{}).Do(request)
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
