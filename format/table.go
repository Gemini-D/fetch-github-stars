package format

import (
	"Gemini-D/fetch-github-stars/repo"
	"github.com/rodaine/table"
)
import "github.com/fatih/color"

func PrintTable(repos []*repo.Repo) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Repo", "Stars", "Forks")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, repo := range repos {
		tbl.AddRow(repo.FullName, repo.StargazersCount, repo.ForksCount)
	}

	tbl.Print()
}
