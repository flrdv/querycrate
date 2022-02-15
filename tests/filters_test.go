package tests

import (
	"testing"

	querycrate "github.com/fakefloordiv/querycrate/qc"
)

func TestNoThirdQuery(t *testing.T) {
	qc := querycrate.NewQueryCrate()

	if err := qc.FromFolder("queries", querycrate.AllowExtensions(".sql")); err != nil {
		t.Fatal("unexpected error during initializing:", err)
	}

	if _, err := qc.Get("query3"); err == nil {
		t.Fatal("query3 shouldn't be here")
	}
}

func TestNoThirdQueryWithDefaultSettings(t *testing.T) {
	qc := querycrate.NewQueryCrate()

	if err := qc.FromFolder("queries"); err != nil {
		t.Fatal("unexpected error during initializing:", err)
	}

	if _, err := qc.Get("query3"); err == nil {
		t.Fatal("query3 shouldn't be here")
	}
}

func TestOnlyTXTFiles(t *testing.T) {
	qc := querycrate.NewQueryCrate()

	if err := qc.FromFolder("queries", querycrate.AllowExtensions(".txt")); err != nil {
		t.Fatal("unexpected error during initializing:", err)
	}

	if _, err := qc.Get("query3"); err != nil {
		t.Fatal("wanted query3 isn't presented")
	}
}

type myCustomFilter struct {
	onlyFilenames []string
}

func (m myCustomFilter) IsAllowed(file querycrate.File) bool {
	for _, allowedFilenames := range m.onlyFilenames {
		if file.Name == allowedFilenames {
			return true
		}
	}

	return false
}

func TestCustomFilter(t *testing.T) {
	onlyQuery1 := myCustomFilter{onlyFilenames: []string{"query1"}}
	qc := querycrate.NewQueryCrate()

	if err := qc.FromFolder("queries", onlyQuery1); err != nil {
		t.Fatal("unexpected error during initializing:", err)
	}

	if _, err := qc.Get("query1"); err != nil {
		t.Fatal("wanted query1 isn't presented")
	}
	if _, err := qc.Get("query2"); err == nil {
		t.Fatal("unexpected query2")
	}
}
