// Scry Info.  All rights reserved.
// license that can be found in the license file.

package VariFlight

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"go.uber.org/zap"

	"github.com/scryinfo/dot/dot"
)

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
		dot.Logger().Debugln("newStorer failed to create table if not exists.", zap.Error(err))
		os.Exit(1)
		return nil
	}
	return &storer{dbx: dbx}
}

func (s *storer) create(data *data) error {
	query := "INSERT INTO flight_data (_token, digest, updated_at_time, flight_data_layout, value_json_string)" +
		" VALUES (:_token, :digest, :updatedAtTime, :flightDataLayout, :valueJSONString)"
	_, err := s.dbx.NamedExec(query, data)
	if err != nil {
		return newDBAccessError(query, err)
	}
	return nil
}

func (s *storer) updateUpdateAtTime(token string, newUpdatedAtTime time.Time) error {
	query := "UPDATE flight_data SET updated_at_time = :updatedAtTime WHERE _token = :_token"
	_, err := s.dbx.NamedExec(query,
		struct {
			token         string
			updatedAtTime time.Time
		}{token, newUpdatedAtTime})
	if err != nil {
		return newDBAccessError(query, err)
	}
	return nil
}

func (s *storer) update(token, newDigest string, newUpdatedAtTime time.Time, valueJSONString string) error {
	query := "UPDATE flight_data SET digest = :digest, updated_at_time = :updatedAtTime, value_json_string = :valueJSONString WHERE _token = :_token"
	_, err := s.dbx.NamedExec(query,
		struct {
			token           string
			digest          string
			updatedAtTime   time.Time
			valueJSONString string
		}{token, newDigest, newUpdatedAtTime, valueJSONString})
	if err != nil {
		return newDBAccessError(query, err)
	}
	return nil
}

// todo: retrieve method for flight data restore.
