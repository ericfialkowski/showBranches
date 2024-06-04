package main

import (
	"errors"
	"flag"
	"github.com/ericfialkowski/env"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/pterm/pterm"
)

var header = []string{"Directory", "Repo", "Main Branch", "Current Branch"}

var diffs bool

func main() {
	defaultBase := env.StringOrDefault("SHOWBRANCHES_DEFAULT", ".")
	flag.BoolVar(&diffs, "d", false, "Only display diffs")

	flag.Parse()
	dirs := flag.Args()

	if len(dirs) == 0 {
		for _, s := range strings.Split(defaultBase, " ") {
			printData(s)
			os.Exit(0)
		}
	}
	for _, base := range dirs {
		printData(base)
	}
}

func printData(base string) {
	data := getBranchInfo(base)
	_ = pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(data).Render()
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
			c, err := r.Config()
			if err != nil {
				panic(err)
			}

			d := make([]string, 4)

			d[0] = dir.Name()
			if dir.Name() == ".git" {
				p, err := filepath.Abs(dir.Name())
				if err != nil {
					panic(err) // shouldn't get an error trying to get full path
				}
				parts := strings.Split(p, string(os.PathSeparator))
				d[0] = parts[len(parts)-2]
			}
			d[1] = c.Remotes["origin"].URLs[0]
			ref, err := r.Reference("refs/remotes/origin/HEAD", false)
			if err != nil {
				if errors.Is(err, plumbing.ErrReferenceNotFound) {
					continue
				}
				panic(err)
			}
			d[2] = last(ref.Target().String(), "/")

			h, err := r.Head()
			if err != nil {
				if errors.Is(err, plumbing.ErrReferenceNotFound) {
					continue
				}
				panic(err)
			}
			d[3] = last(h.Name().String(), "/")
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
