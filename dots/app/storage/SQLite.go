package storage

import (
	"github.com/jinzhu/gorm"
	// blank import for SQLite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/app/storage/definition"
	"go.uber.org/zap"
	"reflect"
	"strconv"
)

// SQLite is an implement of Database, use SQLite
type SQLite struct {
	TableNames     []string
	TableStructure []interface{}

	config sqlConfig
}

type sqlConfig struct {
	DBName string `json:"DBName"`
}

// const
const (
	DatabaseTypeId = "cd947210-6790-4e9f-b73f-63aeba484beb"
	DB             = "sqlite3"
)

// check if 'SQLite' struct implements 'Database' interface
var _ Database = (*SQLite)(nil)

// Create dot.Create
func (s *SQLite) Create(l dot.Line) error {
	s.TableNames = []string{
		"accounts",
		"data_lists",
		"transactions",
		"events",
	}
	s.TableStructure = []interface{}{
		&definition.Account{},
		&definition.DataList{},
		&definition.Transaction{},
		&definition.Event{},
	}

	return nil
}

//construct dot
func newSQLiteDot(conf []byte) (dot.Dot, error) {
	var err error

	dConf := &sqlConfig{}
	err = dot.UnMarshalConfig(conf, dConf)
	if err != nil {
		return nil, err
	}

	d := &SQLite{config: *dConf}

	return d, err
}

// SQLiteTypeLive add a dot component to dot.line with 'line.PreAdd()'
func SQLiteTypeLive() *dot.TypeLives {
	return &dot.TypeLives{
		Meta: dot.Metadata{
			TypeId: DatabaseTypeId,
			NewDoter: func(conf []byte) (dot.Dot, error) {
				return newSQLiteDot(conf)
			},
		},
	}
}

// Start dot.Start
func (s *SQLite) Start(ignore bool) error {
	return s.Init()
}

// Init create tables if not exist
func (s *SQLite) Init() error {
	db, err := gorm.Open(DB, s.config.DBName)
	if err != nil {
		dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in Init()", err))
	}
	defer db.Close()

	for i := range s.TableNames {
		if !db.HasTable(s.TableNames[i]) {
			db.AutoMigrate(s.TableStructure[i])
		}
	}

	return nil
}

// Basic CRUD, support one item and a slice of items

// Insert add record(s) to db
func (s *SQLite) Insert(v interface{}) (int64, error) {
	db, err := gorm.Open(DB, s.config.DBName)
	if err != nil {
		dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in Create()", err))
	}
	defer db.Close()

	value := reflect.ValueOf(v)
	var num int64 = 1
	if value.Kind() != reflect.Slice {
		if err = db.Create(v).Error; err != nil {
			num = 0
		}
	} else {
		var i int
		for i = 0; i < value.Len(); i++ {
			if err = db.Create(value.Index(i).Interface()).Error; err != nil {
				break
			}
		}

		num = int64(i)
	}

	return num, errors.Wrap(err, " error number: "+strconv.FormatInt(num+1, 10)+"th item. ")
}

// Read find records matched from db
func (s *SQLite) Read(out interface{}, order, query string, sql ...interface{}) (int64, error) {
	db, err := gorm.Open(DB, s.config.DBName)
	if err != nil {
		dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in ReadWithOrder()", err))
	}
	defer db.Close()

	t := db.Order(order).Where(query, sql...).Find(out)

	return t.RowsAffected, t.Error
}

// Update update specific item(s) on matched record(s)
func (s *SQLite) Update(out interface{}, m map[string]interface{}, query string, sql ...interface{}) (int64, error) {
	db, err := gorm.Open(DB, s.config.DBName)
	if err != nil {
		dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in Update()", err))
	}
	defer db.Close()

	t := db.Where(query, sql...).Find(out).Updates(m).Find(out)

	return t.RowsAffected, t.Error
}

// Delete remove matched records
func (s *SQLite) Delete(typ interface{}, query string, sql ...interface{}) (int64, error) {
	db, err := gorm.Open(DB, s.config.DBName)
	if err != nil {
		dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in Delete()", err))
	}
	defer db.Close()

	t := db.Where(query, sql...).Delete(typ)

	return t.RowsAffected, t.Error
}

// ReadPage todo: wait client implement
func (s *SQLite) ReadPage() error {
	return nil
}

// UpdateWithoutHooks update without hooks
func (s *SQLite) UpdateWithoutHooks(typ interface{}, m map[string]interface{}, query string, sql ...interface{}) (int64, error) {
	db, err := gorm.Open(DB, s.config.DBName)
	if err != nil {
		dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in UpdateWithoutHooks()", err))
	}
	defer db.Close()

	t := db.Model(typ).Where(query, sql...).Updates(m)

	return t.RowsAffected, t.Error
}
