package tests

import (
	"fmt"
	"testing"

	querycrate "github.com/fakefloordiv/querycrate/qc"
)

func TestNoThirdQuery(t *testing.T) {
	qc, err := querycrate.NewQueryCrate("queries", querycrate.AllowExtensions(".sql"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if _, err = qc.Get("query3"); err == nil {
		t.Fatal("query3 shouldn't be here")
	}
}

func TestNoThirdQueryWithDefaultSettings(t *testing.T) {
	qc, err := querycrate.NewQueryCrate("queries")

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if _, err = qc.Get("query3"); err == nil {
		t.Fatal("query3 shouldn't be here")
	}
}

func TestOnlyTXTFiles(t *testing.T) {
	qc, err := querycrate.NewQueryCrate("queries", querycrate.AllowExtensions(".txt"))

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if _, err = qc.Get("query3"); err != nil {
		t.Fatal("wanted query3 isn't presented")
	}
}

type myCustomFilter struct {
	onlyFilenames []string
}

func (m myCustomFilter) IsAllowed(file querycrate.File) bool {
	fmt.Println("IsAllowed:", file)

	for _, allowedFilenames := range m.onlyFilenames {
		if file.Name == allowedFilenames {
			return true
		}
	}

	return false
}

func TestCustomFilter(t *testing.T) {
	onlyQuery1 := myCustomFilter{onlyFilenames: []string{"query1"}}
	qc, err := querycrate.NewQueryCrate("queries", onlyQuery1)

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if _, err = qc.Get("query1"); err != nil {
		t.Fatal("wanted query1 isn't presented")
	}
	if _, err = qc.Get("query2"); err == nil {
		t.Fatal("unexpected query2")
	}
}
