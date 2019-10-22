package storage

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dp/dots/app/storage/definition"
    "go.uber.org/zap"
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

// CRUD (one item)
func (s *SQLite) Create(v interface{}) error {
    db, err := gorm.Open(DB, DBName)
    if err != nil {
        dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in Create()", err))
    }
    defer db.Close()

    return db.Create(v).Error
}

func (s *SQLite) Read(out interface{}, sql ...interface{}) error {
    db, err := gorm.Open(DB, DBName)
    if err != nil {
        dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in Read()", err))
    }
    defer db.Close()

    return db.First(out, sql...).Error
}

func (s *SQLite) Update(m map[string]interface{}, out interface{}, sql ...interface{}) error {
    db, err := gorm.Open(DB, DBName)
    if err != nil {
        dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in Update()", err))
    }
    defer db.Close()

    return db.First(out, sql...).Updates(m).Error
}

func (s *SQLite) Delete(type_ interface{}, sql ...interface{}) error {
    db, err := gorm.Open(DB, DBName)
    if err != nil {
        dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in Delete()", err))
    }
    defer db.Close()

    return db.Delete(type_, sql...).Error
}

// FindPage, GetSummary
func (s *SQLite) ReadPage() error {
    return nil
}

// FindSome (sorted)
func (s *SQLite) ReadSome(order string, out interface{}, sql ...interface{}) error {
    db, err := gorm.Open(DB, DBName)
    if err != nil {
        dot.Logger().Errorln("Database connect failed. ", zap.NamedError("in ReadSome()", err))
    }
    defer db.Close()

    return db.Order(order).Find(out, sql...).Error
}
