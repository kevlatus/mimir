package mimir

type Scanner interface {
	Scan(dest ...any) error
}

type ScanFunc[T any] func(rows Scanner) (T, error)

func ScanBoolean(scanner Scanner) (bool, error) {
	var isDeleted *bool
	err := scanner.Scan(&isDeleted)
	if err != nil {
		return false, err
	}
	if isDeleted == nil {
		return false, nil
	} else {
		return true, nil
	}
}

func ScanString(scanner Scanner) (string, error) {
	var result string
	err := scanner.Scan(&result)
	return result, err
}

func ScanInt64(scanner Scanner) (int64, error) {
	var record int64
	err := scanner.Scan(&record)
	return record, err
}

func ScanVoid(scanner Scanner) (interface{}, error) {
	return nil, nil
}
