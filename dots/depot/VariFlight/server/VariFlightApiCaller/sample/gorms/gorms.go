package main

import (
	"database/sql"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/db/gorms"
	"github.com/scryinfo/dot/dots/line"
	"github.com/scryinfo/scryg/sutils/ssignal"
	"go.uber.org/zap"
	"log"
	"os"
)

func main()  {
	l, err := line.BuildAndStart(buildNewer)
	if err != nil {
		dot.Logger().Debugln("BuildAndStart failed.", zap.Error(err))
		os.Exit(1)
	}
	defer line.StopAndDestroy(l, false)

	gmdot, err := l.ToInjecter().GetByLiveId(gorms.TypeId)
	if err != nil {
		dot.Logger().Debugln("GetByLiveId(gorms.TypeId) failed.", zap.Error(err))
		os.Exit(1)
	}
	dot.Logger().Debug(func() string {
		return spew.Sprintf("gorms.TypeId: %#+v", gmdot.(*gorms.Gorms))
	})

	gorms := gmdot.(*gorms.Gorms)
	db := gorms.Db.DB()
	driverName := gorms.Db.Dialect().GetName()
	fmt.Printf("db: %#v\n", db)
	fmt.Printf("dialect name: %v\n", driverName)
	testSqlxSqlite3(db, driverName)

	ssignal.WaitCtrlC(nil)

	//testSqlxSqlite3()
}

func buildNewer(l dot.Line) error {
	if err := l.PreAdd(gorms.TypeLives()...); err != nil {
		dot.Logger().Debugln("PreAdd failed.", zap.Error(err))
		os.Exit(1)
	}
	return nil
}

// test db connectivity
func testSqlxSqlite3(db *sql.DB, driverName string)  {
	_ = newStorer(sqlite3(db), driverName)
}

func sqlite3(db *sql.DB) *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", "./testDB.db")
		if err != nil {
			log.Fatalf("failed to open sqlite3. err: %v", err)
		}
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping sqlite3. err: %v", err)
	}
	//defer db.Close()

	fmt.Println("Success to connect sqlite3")
	return db
}

var schema = `
	CREATE TABLE IF NOT EXISTS flight_data (
		_token text,
		digest text,
		updated_at_time text,
		flight_data_layout text,
		value_json_string text);`

type storer struct {
	dbx *sqlx.DB
}

func newStorer(db *sql.DB, driverName string) *storer {
	dbx := sqlx.NewDb(db, driverName)
	_, err := dbx.Exec(schema)
	if err != nil {
		log.Fatalf("failed exec schema. err: %v", err)
	}
	fmt.Println("success to exec schema")
	return &storer{dbx: dbx}
}