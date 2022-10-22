package mimir

type ArgsFunc[T any] func(record T) []any

type SelectAllQuery[T any, K string | int64, R any] interface {
	SelectAll(ex QueryExecutor, args ...any) ([]R, error)
}

func (q *query[T, K, R]) SelectAll(
	ex QueryExecutor,
	args ...any,
) ([]R, error) {
	return q.QueryRows(ex, args...)
}

func NewSelectAllQuery[T any, K string | int64, R any](
	statement string,
	scanFunc ScanFunc[R],
) SelectAllQuery[T, K, R] {
	return NewQuery[T, K](statement, scanFunc, nil)
}

type SelectByIdQuery[T any, K string | int64, R any] interface {
	SelectById(ex QueryExecutor, id K, args ...any) (R, error)
}

func (q *query[T, K, R]) SelectById(
	ex QueryExecutor,
	id K,
	args ...any,
) (R, error) {
	allArgs := []any{id}
	allArgs = append(allArgs, args...)
	return q.QuerySingleRow(ex, allArgs...)
}

func NewSelectByIdQuery[T any, K string | int64, R any](
	statement string,
	scanFunc ScanFunc[R],
) SelectByIdQuery[T, K, R] {
	return NewQuery[T, K](statement, scanFunc, nil)
}

func (q *query[T, K, R]) Insert(
	ex QueryExecutor,
	entity T,
	args ...any,
) (R, error) {
	allArgs := []any{}
	allArgs = append(allArgs, args...)
	allArgs = append(allArgs, q.getRecordArgs(entity)...)
	return q.QuerySingleRow(
		ex,
		allArgs...,
	)
}

type InsertQuery[T any, K string | int64, R any] interface {
	Insert(ex QueryExecutor, entity T, args ...any) (R, error)
}

func NewInsertQuery[T any, K string | int64, R any](
	statement string,
	scanFunc ScanFunc[R],
	getEntityArgs ArgsFunc[T],
) InsertQuery[T, K, R] {
	return NewQuery[T, K](statement, scanFunc, getEntityArgs)
}

type UpdateQuery[T any, K string | int64, R any] interface {
	Update(ex QueryExecutor, entity T, args ...any) (R, error)
}

func (q *query[T, K, R]) Update(
	ex QueryExecutor,
	entity T,
	args ...any,
) (R, error) {
	allArgs := []any{}
	allArgs = append(allArgs, args...)
	allArgs = append(allArgs, q.getRecordArgs(entity)...)
	return q.QuerySingleRow(ex, allArgs...)
}

func NewUpdateQuery[T any, K string | int64, R any](
	statement string,
	scanFunc ScanFunc[R],
	getEntityArgs ArgsFunc[T],
) UpdateQuery[T, K, R] {
	return NewQuery[T, K](statement, scanFunc, getEntityArgs)
}

type DeleteQuery[T any, K string | int64, R any] interface {
	Query[T, K, R]

	DeleteById(ex QueryExecutor, id K, args ...any) (R, error)
}

func (q *query[T, K, R]) DeleteById(
	ex QueryExecutor,
	id K,
	args ...any,
) (R, error) {
	allArgs := []any{}
	allArgs = append(allArgs, args...)
	allArgs = append(allArgs, id)
	return q.QuerySingleRow(ex, allArgs...)
}

func NewDeleteQuery[T any, K string | int64, R any](
	statement string,
	scanFunc ScanFunc[R],
) DeleteQuery[T, K, R] {
	return NewQuery[T, K](statement, scanFunc, nil)
}
