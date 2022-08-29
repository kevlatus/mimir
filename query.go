package mimir

import "database/sql"

func querySingleRow[T any](
	db *sql.DB,
	scanner ScanFunc[T],
	queryText string,
	args ...any) (T, error) {
	var result T

	row := db.QueryRow(queryText, args...)
	var err error
	result, err = scanner(row)

	return result, err
}

func queryRows[T any](
	db *sql.DB,
	scanner ScanFunc[T],
	queryText string,
	args ...any,
) ([]T, error) {
	result := []T{}

	rows, err := db.Query(queryText, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		record, err := scanner(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}

	return result, err
}

type Query[T any, K string | int64, R any] interface {
	QuerySingleRow(database *Database, args ...any) (R, error)

	QueryRows(database *Database, args ...any) ([]R, error)
}

type query[T any, K string | int64, R any] struct {
	statement     string
	scanFunc      ScanFunc[R]
	getRecordArgs ArgsFunc[T]
}

func NewQuery[T any, K string | int64, R any](
	statement string,
	scanFunc ScanFunc[R],
	getRecordArgs ArgsFunc[T],
) *query[T, K, R] {
	return &query[T, K, R]{
		statement:     statement,
		scanFunc:      scanFunc,
		getRecordArgs: getRecordArgs,
	}
}

func (q *query[T, K, R]) QuerySingleRow(
	database *Database,
	args ...any,
) (R, error) {
	var result R
	err := database.WithConn(func(db *sql.DB) error {
		var err error
		result, err = querySingleRow(db, q.scanFunc, q.statement, args...)
		return err
	})
	return result, err
}

func (q *query[T, K, R]) QueryRows(
	database *Database,
	args ...any,
) ([]R, error) {
	var result []R
	err := database.WithConn(func(db *sql.DB) error {
		var err error
		result, err = queryRows(db, q.scanFunc, q.statement, args...)
		return err
	})
	return result, err
}
