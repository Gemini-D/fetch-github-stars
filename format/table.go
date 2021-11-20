package format

import (
	"Gemini-D/fetch-github-stars/repo"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"sort"
)

func PrintTable(repos []*repo.Repo) {
	sort.SliceStable(repos, func(i, j int) bool {
		if repos[i].StargazersCount > repos[j].StargazersCount {
			return true
		}
		return false
	})

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Repo", "Stars", "Forks")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, repo := range repos {
		tbl.AddRow(repo.FullName, repo.StargazersCount, repo.ForksCount)
	}

	tbl.Print()
}
