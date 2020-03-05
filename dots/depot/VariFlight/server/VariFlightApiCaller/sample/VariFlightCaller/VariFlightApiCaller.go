package main

import (
	"database/sql"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/db/gorms"
	"github.com/scryinfo/dot/dots/line"
	VariFlight "github.com/scryinfo/dp/dots/depot/VariFlight/server/VariFlightApiCaller"
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
		return spew.Sprintf("gorms: %v", gmdot.(*gorms.Gorms))
	})

	vfdot, err := l.ToInjecter().GetByLiveId(VariFlight.VariFlightApiCallerTypeId)
	if err != nil {
		dot.Logger().Debugln("GetByLiveId(VariFlightApiCallerTypeId) failed.", zap.Error(err))
		os.Exit(1)
	}
	dot.Logger().Debug(func() string {
		return spew.Sprintf("VariFlightApiCaller: %v", vfdot.(*VariFlight.VariFlightApiCaller))
	})

	fmt.Println("VariFlightApiCaller now can work normally!")
	ssignal.WaitCtrlC(nil)

	//testSqlxSqlite3()

}

func buildNewer(l dot.Line) error {
	if err := l.PreAdd(VariFlight.VariFlightCallerTypeLives()...); err != nil {
		dot.Logger().Debugln("PreAdd failed.", zap.Error(err))
		os.Exit(1)
	}
	return nil
}

// test db connectivity
func testSqlxSqlite3()  {
	_ = newStorer(sqlite3(), "sqlite3")
}

func sqlite3() *sql.DB {
	db, err := sql.Open("sqlite3", "./testDB.db")
	if err != nil {
		log.Fatalf("failed to open sqlite3. err: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping sqlite3. err: %v", err)
	}
	//defer db.Close()

	fmt.Println("Success to open sqlite3")
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