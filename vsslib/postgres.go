package vsslib

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func NewPostgresSession(dbHost string, dbPort string, dbUser string, dbPass string, dbName string) (*sql.DB, error) {
	var err error
	var db *sql.DB

	dbSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err = sql.Open("postgres", dbSource)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	var openConnections int = 25
	openConnectionsValue, openConnectionsPresent := os.LookupEnv("DB_MAX_OPEN_CONNECTIONS")
	if openConnectionsPresent {
		openConnections, err = strconv.Atoi(openConnectionsValue)
		if err != nil {
			return nil, err
		}
	}
	db.SetMaxOpenConns(openConnections)

	var idleConnection int = 25
	idleConnectionValue, idleConnectionPresent := os.LookupEnv("DB_MAX_IDLE_CONNECTIONS")
	if idleConnectionPresent {
		idleConnection, err = strconv.Atoi(idleConnectionValue)
		if err != nil {
			return nil, err
		}
	}
	db.SetMaxIdleConns(idleConnection)

	return db, nil
}
