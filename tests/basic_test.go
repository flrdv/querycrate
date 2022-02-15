package tests

import (
	"errors"
	"fmt"
	"testing"

	querycrate "github.com/fakefloordiv/querycrate/qc"
)

func wantQuery(qc querycrate.QueryCrate, queryName, wantQuery string) error {
	query, err := qc.Get(queryName)

	if err != nil {
		return err
	}

	if query != wantQuery {
		return errors.New(fmt.Sprintf(`unexpected query: wanted "%s", got "%s"\n`,
			wantQuery, query))
	}

	return nil
}

func TestQueriesFromRoot(t *testing.T) {
	qc := querycrate.NewQueryCrate()

	if err := qc.FromFolder("queries"); err != nil {
		t.Fatal("unexpected error during initializing:", err)
	}

	query1 := "query1"
	query2 := "query2"

	if err := wantQuery(qc, query1, query1); err != nil {
		t.Fatal(err)
	}
	if err := wantQuery(qc, query2, query2); err != nil {
		t.Fatal(err)
	}
}

func TestQueriesFromSomeRepository(t *testing.T) {
	qc := querycrate.NewQueryCrate()

	if err := qc.FromFolder("queries"); err != nil {
		t.Fatal("unexpected error during initializing:", err)
	}

	query1 := "somerepository/query1"
	query2 := "somerepository/query2"

	if err := wantQuery(qc, query1, query1); err != nil {
		t.Fatal(err)
	}
	if err := wantQuery(qc, query2, query2); err != nil {
		t.Fatal(err)
	}
}

func TestQueriesFromSubRepository(t *testing.T) {
	qc := querycrate.NewQueryCrate()

	if err := qc.FromFolder("queries"); err != nil {
		t.Fatal("unexpected error during initializing:", err)
	}

	query1 := "somerepository/subrepository/query1"
	query2 := "somerepository/subrepository/query2"

	if err := wantQuery(qc, query1, query1); err != nil {
		t.Fatal(err)
	}
	if err := wantQuery(qc, query2, query2); err != nil {
		t.Fatal(err)
	}
}

func TestAddQuery(t *testing.T) {
	qc := querycrate.NewQueryCrate()

	if err := qc.AddQuery("queries/query1.sql"); err != nil {
		t.Fatal("unexpected error while adding queries/query1.sql:", err)
	}

	err := wantQuery(qc, "queries/query1", "query1")

	if err != nil {
		t.Fatal("unexpected error while getting query:", err)
	}
}
