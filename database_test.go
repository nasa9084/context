package context_test

import (
	"database/sql"
	"testing"

	context "github.com/nasa9084/subcontext"
)

func TestDatabase(t *testing.T) {
	db := &sql.DB{}
	candidates := []struct {
		name string
		in     context.Context
		expect *sql.DB
	}{
		{
			"empty context",
			context.Background(),
			nil,
		}, {
			"WithDatabase",
			context.WithDatabase(context.Background(), db),
			db,
		},
	}
	for _, c := range candidates {
		t.Log(c.name)
		db := context.Database(c.in)
		if db != c.expect {
			t.Errorf(`"%v" != "%v"`, db, c.expect)
			return
		}
	}
}
