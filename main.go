package main

import (
	"errors"
	"flag"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/pterm/pterm"
)

var header = []string{"Repo", "Main Branch", "Current Branch"}

var diffs bool

func main() {
	flag.BoolVar(&diffs, "d", false, "Only display diffs")

	flag.Parse()
	dirs := flag.Args()

	if len(dirs) == 0 {
		printData(".")
		os.Exit(0)
	}
	for _, base := range dirs {
		printData(base)
	}
}

func printData(base string) {
	data := getBranchInfo(base)
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(data).Render()
}

func getBranchInfo(base string) [][]string {
	dirs, err := os.ReadDir(base)
	if err != nil {
		panic(err)
	}
	var data [][]string
	data = append(data, header)
	for _, dir := range dirs {
		if dir.IsDir() {
			r, err := git.PlainOpen(path.Join(base, dir.Name()))
			if err != nil {
				if errors.Is(err, git.ErrRepositoryNotExists) {
					continue
				}
				panic(err)
			}
			d := make([]string, 3)

			d[0] = dir.Name()
			ref, err := r.Reference("refs/remotes/origin/HEAD", false)
			if err != nil {
				if errors.Is(err, plumbing.ErrReferenceNotFound) {
					continue
				}
				panic(err)
			}
			d[1] = last(ref.Target().String(), "/")

			h, err := r.Head()
			if err != nil {
				if errors.Is(err, plumbing.ErrReferenceNotFound) {
					continue
				}
				panic(err)
			}
			d[2] = last(h.Name().String(), "/")
			if !diffs || !strings.EqualFold(d[1], d[2]) {
				data = append(data, d)
			}
		}
	}
	return data
}

func last(s, sep string) string {
	splits := strings.Split(s, sep)
	return splits[len(splits)-1]
}
