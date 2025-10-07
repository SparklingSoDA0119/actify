package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	pDb  *sql.DB
	isConnected bool
}


func NewPostgresDb() *PostgresDB {
	return &PostgresDB{isConnected: false}
}

func (db *PostgresDB) InitializePostgres(connStr string) error {
	pDb, err := connectPostgres(connStr)
	if err != nil {
		return fmt.Errorf("database connection error(err: %v)", err)
	}

	db.pDb = pDb

	err = db.checkTable()
	if err != nil {
		return fmt.Errorf("check database table error(err: %v)", err)
	}

	db.isConnected = true
	return nil
}


func connectPostgres(connStr string) (*sql.DB, error) {
	pDb, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = pDb.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Info: Postgress Database connection success")

	return pDb, nil
}


func (pDb *PostgresDB) Destroy() {
	defer pDb.pDb.Close()
	pDb.isConnected = false
}


func (pDb *PostgresDB) checkTable() error {
	
	return nil
}