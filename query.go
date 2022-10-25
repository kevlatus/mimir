package mimir

import "database/sql"

func querySingleRow[T any](
	ex QueryExecutor,
	scanner ScanFunc[T],
	queryText string,
	args ...any) (T, error) {
	var result T

	row := ex.QueryRow(queryText, args...)
	var err error
	result, err = scanner(row)

	return result, err
}

func queryRows[T any](
	ex QueryExecutor,
	scanner ScanFunc[T],
	queryText string,
	args ...any,
) ([]T, error) {
	result := []T{}

	rows, err := ex.Query(queryText, args...)
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
	Exec(ex QueryExecutor, args ...any) (sql.Result, error)

	QuerySingleRow(ex QueryExecutor, args ...any) (R, error)

	QueryRows(ex QueryExecutor, args ...any) ([]R, error)
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

func (q *query[T, K, R]) Exec(
	ex QueryExecutor,
	args ...any,
) (sql.Result, error) {
	return ex.Exec(q.statement, args...)
}

func (q *query[T, K, R]) QuerySingleRow(
	ex QueryExecutor,
	args ...any,
) (R, error) {
	return querySingleRow(ex, q.scanFunc, q.statement, args...)
}

func (q *query[T, K, R]) QueryRows(
	ex QueryExecutor,
	args ...any,
) ([]R, error) {
	return queryRows(ex, q.scanFunc, q.statement, args...)
}
