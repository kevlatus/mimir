package mimir

type ArgsFunc[T any] func(record T) []any

type SelectAllQuery[T any, K string | int64, R any] interface {
	SelectAll(database *Database, args ...any) ([]R, error)
}

func (q *query[T, K, R]) SelectAll(
	database *Database,
	args ...any,
) ([]R, error) {
	return q.QueryRows(database, args...)
}

func NewSelectAllQuery[T any, K string | int64, R any](
	statement string,
	scanFunc ScanFunc[R],
) SelectAllQuery[T, K, R] {
	return NewQuery[T, K](statement, scanFunc, nil)
}

type SelectByIdQuery[T any, K string | int64, R any] interface {
	SelectById(database *Database, id K, args ...any) (R, error)
}

func (q *query[T, K, R]) SelectById(
	database *Database,
	id K,
	args ...any,
) (R, error) {
	allArgs := []any{id}
	allArgs = append(allArgs, args...)
	return q.QuerySingleRow(database, allArgs...)
}

func NewSelectByIdQuery[T any, K string | int64, R any](
	statement string,
	scanFunc ScanFunc[R],
) SelectByIdQuery[T, K, R] {
	return NewQuery[T, K](statement, scanFunc, nil)
}

func (q *query[T, K, R]) Insert(
	database *Database,
	entity T,
	args ...any,
) (R, error) {
	allArgs := []any{}
	allArgs = append(allArgs, args...)
	allArgs = append(allArgs, q.getRecordArgs(entity)...)
	return q.QuerySingleRow(
		database,
		allArgs...,
	)
}

type InsertQuery[T any, K string | int64, R any] interface {
	Insert(database *Database, entity T, args ...any) (R, error)
}

func NewInsertQuery[T any, K string | int64, R any](
	statement string,
	scanFunc ScanFunc[R],
	getEntityArgs ArgsFunc[T],
) InsertQuery[T, K, R] {
	return NewQuery[T, K](statement, scanFunc, getEntityArgs)
}

type UpdateQuery[T any, K string | int64, R any] interface {
	Update(database *Database, entity T, args ...any) (R, error)
}

func (q *query[T, K, R]) Update(
	database *Database,
	entity T,
	args ...any,
) (R, error) {
	allArgs := []any{}
	allArgs = append(allArgs, args...)
	allArgs = append(allArgs, q.getRecordArgs(entity)...)
	return q.QuerySingleRow(database, allArgs...)
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

	DeleteById(database *Database, id K, args ...any) (R, error)
}

func (q *query[T, K, R]) DeleteById(
	database *Database,
	id K,
	args ...any,
) (R, error) {
	allArgs := []any{}
	allArgs = append(allArgs, args...)
	allArgs = append(allArgs, id)
	return q.QuerySingleRow(database, allArgs...)
}

func NewDeleter[T any, K string | int64, R any](
	statement string,
	scanFunc ScanFunc[R],
) DeleteQuery[T, K, R] {
	return NewQuery[T, K](statement, scanFunc, nil)
}
