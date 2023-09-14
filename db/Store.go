package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pwsdc/web-mud/arg"
	"github.com/pwsdc/web-mud/db/dbg"
	"github.com/pwsdc/web-mud/shared"
)

type tStore struct {
	shared.HasResults
	conn  *sql.DB
	Query *dbg.Queries
}

var Store *tStore

func init() {
	Store = &tStore{
		Query: nil,
	}
	shared.HasResultsInit(Store)
}

func (store *tStore) Disconnect() {
	store.Query = nil
	if store.conn == nil {
		return
	}
	err := store.conn.Ping()
	if err != nil {
		store.conn = nil
		return
	}
	store.Log("Disconnected from database.")
	store.conn.Close()
}

func (store *tStore) Connect() error {
	store.Disconnect()
	cnx := ""
	if arg.Config.Http.DockerMode() {
		cnx = arg.Config.Db.ConnectString()
	} else {
		cnx = arg.Config.Db.ConnectStringLocalhost()
	}
	postgres, err := sql.Open("postgres", cnx)
	if err != nil {
		store.Error(fmt.Sprintf("Failed to connect to %s.\n\tError:%s", cnx, err.Error()))
		return err
	}
	err = postgres.Ping()
	if err != nil {
		store.Error(fmt.Sprintf("Ping failed on new postgres connection to %s. \n\tError:%s", cnx, err.Error()))
		return err
	} else {
		store.Log("Pinged postgres.")
	}
	store.conn = postgres
	store.Query = dbg.New(store.conn)
	store.Log("Succesfully connected to postgres.")
	return nil
}

func (store *tStore) PrintLogs() {
	logs := *store.GetLogs()
	for _, log := range logs {
		fmt.Println("db log:", log)
	}
	errs := *store.GetErrors()
	for _, err := range errs {
		fmt.Println("db err:", err)
	}
}
