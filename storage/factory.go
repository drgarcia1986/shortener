package storage

const (
	FakeType = iota
	SQLiteType
)

func New(t int, path string) Storage {
	switch t {
	case FakeType:
		return NewFake()
	case SQLiteType:
		return NewSQLite(path)
	default:
		return NewFake()
	}
}
