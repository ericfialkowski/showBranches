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

const (
	DIR_HEADER  = iota
	REPO_HEADER = iota
	MAIN_HEADER = iota
	CURR_HEADER = iota
)

var diffs bool

func main() {
	defaultBase := env.StringOrDefault("SHOWBRANCHES_DEFAULT", ".")
	flag.BoolVar(&diffs, "d", false, "Only display diffs")

	flag.Parse()
	dirs := flag.Args()

	if len(dirs) == 0 {
		for _, s := range strings.Split(defaultBase, " ") {
			printData(s)
		}
		os.Exit(0)
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

			d[DIR_HEADER] = dir.Name()
			if dir.Name() == ".git" {
				p, err := filepath.Abs(dir.Name())
				if err != nil {
					panic(err) // shouldn't get an error trying to get full path
				}
				parts := strings.Split(p, string(os.PathSeparator))
				d[DIR_HEADER] = parts[len(parts)-2]
			}
			if c.Remotes["origin"] != nil && len(c.Remotes["origin"].URLs) > 0 {
				d[REPO_HEADER] = c.Remotes["origin"].URLs[0]
			} else {
				d[REPO_HEADER] = "<no remote>"
			}
			ref, err := r.Reference("refs/remotes/origin/HEAD", false)
			if err != nil {
				if errors.Is(err, plumbing.ErrReferenceNotFound) {
					d[MAIN_HEADER] = "<no remote>"
				} else {
					panic(err)
				}
			} else {
				d[MAIN_HEADER] = last(ref.Target().String(), "/")
			}

			h, err := r.Head()
			if err != nil {
				if errors.Is(err, plumbing.ErrReferenceNotFound) {
					d[CURR_HEADER] = "<no remote>"
				} else {
					panic(err)
				}
			} else {
				d[CURR_HEADER] = last(h.Name().String(), "/")
			}

			if !diffs || !strings.EqualFold(d[MAIN_HEADER], d[CURR_HEADER]) {
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
