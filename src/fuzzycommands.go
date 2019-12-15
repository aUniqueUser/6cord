package main

import (
	"sort"
	"strings"

	"github.com/diamondburned/tview/v2"
	"github.com/sahilm/fuzzy"
)

// String returns the fuzzy search part of the struct
func (cmds Commands) String(i int) string {
	return cmds[i].Command
}

// Len returns the length of the Emojis slice
func (cmds Commands) Len() int {
	return len(cmds)
}

func fuzzyCommands(last string) {
	var fuzzied Commands

	if len(last) > 1 {
		results := fuzzy.FindFrom(
			strings.TrimPrefix(last, "/"),
			commands,
		)

		for i, r := range results {
			if i == 10 {
				break
			}

			fuzzied = append(
				fuzzied,
				commands[r.Index],
			)
		}

		sort.Slice(fuzzied, func(i, j int) bool {
			return len(fuzzied[i].Command) < len(fuzzied[j].Command)
		})
	} else {
		fuzzied = append(fuzzied, commands...)
	}

	clearList()

	if len(fuzzied) > 0 {
		for i, u := range fuzzied {
			autocomp.InsertItem(i, &tview.ListItem{
				"[::b]" + u.Command + "[::-] - " + tview.Escape(u.Description),
				"", 0, nil,
			})
		}

		rightflex.ResizeItem(autocomp, min(len(fuzzied), 10), 1)

		autofillfunc = func(i int) {
			input.SetText(fuzzied[i].Command + " ")
			clearList()
			app.SetFocus(input)
		}

	} else {
		rightflex.ResizeItem(autocomp, 1, 1)
	}

	app.Draw()
}
