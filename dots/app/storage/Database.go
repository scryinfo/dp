package storage

// Database define functions a db instance need to implement
type Database interface {
	// connect and init
	Init() error

	// Basic CRUD

	Insert(v interface{}) (int64, error)
	Read(out interface{}, order, query string, sql ...interface{}) (int64, error)
	Update(out interface{}, m map[string]interface{}, query string, sql ...interface{}) (int64, error)
	Delete(typ interface{}, query string, sql ...interface{}) (int64, error)

	// Customized CRUD

	// ReadPage, todo
	ReadPage() error

	// Update one item with hooks
	UpdateWithoutHooks(typ interface{}, m map[string]interface{}, query string, sql ...interface{}) (int64, error)
}
