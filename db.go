package mimir

import "database/sql"

type Database struct {
	connStr    string
	driverName string
}

func NewDatabase(driverName string, connStr string) *Database {
	return &Database{connStr: connStr, driverName: driverName}
}

func (d *Database) WithConn(f func(db *sql.DB) error) error {
	db, err := sql.Open(d.driverName, d.connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	return f(db)
}

type QueryExecutor interface {
	Query(query string, args ...any) (*sql.Rows, error)

	QueryRow(query string, args ...any) *sql.Row

	Exec(query string, args ...any) (sql.Result, error)
}
