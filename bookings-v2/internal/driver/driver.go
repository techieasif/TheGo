package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDBConn = 10
const maxIdleDBConn = 5
const maxDbLifeTime = time.Minute * 5

// ConnectSQL connectSQL creates connection to database
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)

	if err != nil {
		panic(err)
	}

	d.SetConnMaxIdleTime(maxIdleDBConn)
	d.SetMaxOpenConns(maxOpenDBConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	dbConn.SQL = d

	if err = testDBConnection(d); err != nil {
		return nil, err
	}

	return dbConn, nil

}

// testDBConnection tests the created connection
func testDBConnection(d *sql.DB) error {
	if err := d.Ping(); err != nil {
		return err
	}
	return nil
}

// NewDatabase opens a new database instance.
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
