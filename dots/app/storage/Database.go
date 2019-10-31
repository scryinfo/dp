package storage

type Database interface {
	// connect and init
	Init() error

	// Basic CRUD

	Insert(v interface{}) (int64, error)
	Read(out interface{}, order, query string, sql ...interface{}) (int64, error)
	Update(out interface{}, m map[string]interface{}, query string, sql ...interface{}) (int64, error)
	Delete(type_ interface{}, query string, sql ...interface{}) (int64, error)

	// Customized CRUD

	// ReadPage, todo: wait js implement
	ReadPage() error

	// Update one item with hooks
	UpdateWithoutHooks(type_ interface{}, m map[string]interface{}, query string, sql ...interface{}) (int64, error)
}
