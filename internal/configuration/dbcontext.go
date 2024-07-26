package configuration

import (
	"database/sql"
	"encoding/json"
)

type DBContext map[string]func() *sql.DB

func (d DBContext) Get(name string) *sql.DB {
	if m, ok := d[name]; ok {
		return m()
	}
	return nil
}

func (d DBContext) Set(name string, fn func() *sql.DB) {
	d[name] = fn
}

func builder(config DatabaseConf) func() *sql.DB {
	once := make(chan int, 1)
	once <- 0

	var conn *sql.DB
	return func() *sql.DB {
		entry := <-once
		defer func() {
			once <- entry + 1
		}()

		if entry == 0 {
			conn = sql.OpenDB(config.toParams())
			// conn.

		}
		return conn
	}
}

func New() DBContext {
	var databases []DatabaseConf
	if err := json.Unmarshal([]byte(databaseFile), &databases); err != nil {
		panic(err)
	}

	db := make(DBContext)
	for _, conf := range databases {
		db.Set(conf.Name, builder(conf.Parse()))
	}
	return db
}
