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
