package storage

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/app/storage/definition"
    "go.uber.org/zap"
    "reflect"
    "strconv"
)

type SQLite struct {
}

// todo: configuration
const (
    DB = "sqlite3"
    DBName = "app.db"
)

var (
    tableName = []string{
        "accounts",
        "data_lists",
        "transactions",
        "events",
    }
    tableStructure = []interface{}{
        &definition.Account{},
        &definition.DataList{},
        &definition.Transaction{},
        &definition.Event{},
    }
)

// check if 'SQLite' struct implements 'Database' interface
var _ Database = (*SQLite)(nil)

// create tables if not exist
func (s *SQLite) Init() error {
    db, err := gorm.Open(DB, DBName)
    if err != nil {
        dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in Init()", err))
    }
    defer db.Close()

    for i := range tableName {
        if !db.HasTable(tableName[i]) {
            db.AutoMigrate(tableStructure[i])
        }
    }

    return nil
}

// Basic CRUD
func (s *SQLite) Create(v interface{}) (int64, error) {
    db, err := gorm.Open(DB, DBName)
    if err != nil {
        dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in Create()", err))
    }
    defer db.Close()

    value := reflect.ValueOf(v)
    num := 1
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

        num = i
    }

    return int64(num), errors.Wrap(err, " error number: " + strconv.Itoa(num) + "th item. ")
}

func (s *SQLite) Read(out interface{}, order, query string, sql ...interface{}) (int64, error) {
    db, err := gorm.Open(DB, DBName)
    if err != nil {
        dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in ReadWithOrder()", err))
    }
    defer db.Close()

    t := db.Order(order).Where(query, sql).Find(out)

    return t.RowsAffected, t.Error
}

func (s *SQLite) Update(out interface{}, m map[string]interface{}, query string, sql ...interface{}) (int64, error) {
    db, err := gorm.Open(DB, DBName)
    if err != nil {
        dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in UpdateOneItemWithHooks()", err))
    }
    defer db.Close()

    t := db.Where(query, sql).Find(out).Updates(m)

    return t.RowsAffected, t.Error
}

func (s *SQLite) Delete(type_ interface{}, query string, sql ...interface{}) (int64, error) {
    db, err := gorm.Open(DB, DBName)
    if err != nil {
        dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in Delete()", err))
    }
    defer db.Close()

    t := db.Where(query, sql).Delete(type_)

    return t.RowsAffected, t.Error
}

// todo: wait js implement
func (s *SQLite) ReadPage() error {
    return nil
}

// update without hooks
func (s *SQLite) UpdateWithoutHooks(type_ interface{}, m map[string]interface{}, query string, sql ...interface{}) (int64, error) {
    db, err := gorm.Open(DB, DBName)
    if err != nil {
        dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in Update()", err))
    }
    defer db.Close()

    t := db.Model(type_).Where(query, sql).Updates(m)

    return t.RowsAffected, t.Error
}
