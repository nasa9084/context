package context

import (
	"database/sql"
	"fmt"
)

type databaseCtx struct {
	Context
	database *sql.DB
}

// WithDatabase returns a copy of parent in which the database
// associated with.
func WithDatabase(parent Context, database *sql.DB) Context {
	if database == nil {
		panic("nil database")
	}
	return &databaseCtx{parent, database}
}

func (c *databaseCtx) String() string {
	return fmt.Sprintf("%v.WithDatabase(%#v)", c.Context, c.database)
}

func (c *databaseCtx) Database() *sql.DB {
	if c.database != nil {
		return c.database
	}
	return nil
}

// Database extracts *sql.DB from Context.
// This will return nil when no *sql.DB is associated with.
func Database(ctx Context) *sql.DB {
	if dbCtx, ok := ctx.(*databaseCtx); ok {
		if dbCtx.Database() != nil {
			return dbCtx.Database()
		}
		return Database(dbCtx.Context)
	}
	return nil
}
