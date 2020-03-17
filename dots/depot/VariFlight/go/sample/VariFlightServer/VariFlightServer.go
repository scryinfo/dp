package main

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dot/dots/db/gorms"
	"github.com/scryinfo/dot/dots/gindot"
	"github.com/scryinfo/dot/dots/grpc/gserver"
	"github.com/scryinfo/dot/dots/line"
	"github.com/scryinfo/dp/dots/depot/VariFlight/go"
	vfApiCaller "github.com/scryinfo/dp/dots/depot/VariFlight/go/VariFlightApiCaller"
	"github.com/scryinfo/scryg/sutils/ssignal"
	"go.uber.org/zap"
	"log"
	"os"
	"reflect"
)

func main()  {
	// build line and configure dots
	l, err := line.BuildAndStart(buildNewer)
	if err != nil {
		dot.Logger().Debugln("BuildAndStart failed.", zap.Error(err))
		os.Exit(1)
	}
	defer line.StopAndDestroy(l, false)

	// get gorms.Gorms component
	gmdot, err := l.ToInjecter().GetByLiveId(gorms.TypeId)
	if err != nil {
		dot.Logger().Debugln("GetByLiveId(gorms.TypeId) failed.", zap.Error(err))
		os.Exit(1)
	}
	dot.Logger().Debug(func() string {
		return fmt.Sprintf("gorms: %#+v", gmdot.(*gorms.Gorms))
	})
	fmt.Println("Gorms component now can work normally!")

	// get VariFlightApiCaller.VariFlightApiCaller component
	vfdot, err := l.ToInjecter().GetByLiveId(vfApiCaller.VariFlightApiCallerTypeId)
	if err != nil {
		dot.Logger().Debugln("GetByLiveId(VariFlightApiCallerTypeId) failed.", zap.Error(err))
		os.Exit(1)
	}
	dot.Logger().Debug(func() string {
		return fmt.Sprintf("VariFlightApiCaller: %#+v", vfdot.(*vfApiCaller.VariFlightApiCaller))
	})

	fmt.Println("VariFlightApiCaller component now can work normally!")

	// get gindot.Engine component
	engine, err := l.ToInjecter().GetByLiveId(gindot.EngineTypeId)
	if err != nil {
		dot.Logger().Debugln("GetByliveId(gindot.EngineTypeId) failed.", zap.Error(err))
		os.Exit(1)
	}
	dot.Logger().Debug(func() string {
		return fmt.Sprintf("Engine: %#+v", engine.(*gindot.Engine))
	})
	fmt.Println("Engine component now can work normally.")

	// get WebSocket component
	ws, err := l.ToInjecter().GetByLiveId(gserver.WebSocketTypeId)
	if err != nil {
		dot.Logger().Debugln("GetByliveId(gserver.WebSocketTypeId) failed.", zap.Error(err))
		os.Exit(1)
	}
	dot.Logger().Debug(func() string {
		return fmt.Sprintf("WebSocket: %#+v", ws.(*gserver.WebSocket))
	})
	fmt.Println("WebSocket component now can work normally.")

	//get VariFlightServer component
	vfServer, err := l.ToInjecter().GetByType(reflect.TypeOf((*_go.VariFlightServer)(nil)))
	//vfServer, err := l.ToInjecter().GetByLiveId(go.VariFlightServerTypeId)
	if err != nil {
		dot.Logger().Debugln("GetByliveId(go.VariFlightServerTypeId) failed.", zap.Error(err))
		os.Exit(1)
	}
	dot.Logger().Debug(func() string {
		return fmt.Sprintf("VariFlightServer: %#+v", vfServer.(*_go.VariFlightServer))
	})
	fmt.Println("VariFlightServer component now can work normally.")

	//testSqlxSqlite3()

	ssignal.WaitCtrlC(nil)
}

func buildNewer(l dot.Line) error {
	if err := l.PreAdd(_go.VariFlightServerTypeLives()...); err != nil {
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
		log.Fatalf("failed to open sqlite3. err: %#+v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping sqlite3. err: %#+v", err)
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
		log.Fatalf("failed exec schema. err: %#+v", err)
	}
	fmt.Println("success to exec schema")
	return &storer{dbx: dbx}
}