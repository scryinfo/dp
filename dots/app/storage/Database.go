package storage

type Database interface {
	// connect
	Init() error

	// CRUD (one item)
	Create(v interface{}) error
	Read(out interface{}, sql ...interface{}) error
	Update(m map[string]interface{}, out interface{}, sql ...interface{}) error
	Delete(type_ interface{}, sql ...interface{}) error

	// FindPage, todo: js not implement, wait
	ReadPage() error

	// FindSome (sorted)
	ReadSome(order string, out interface{}, sql ...interface{}) error
}
