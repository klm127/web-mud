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

func (self *tStore) Disconnect() {
	self.Query = nil
	if self.conn == nil {
		return
	}
	err := self.conn.Ping()
	if err != nil {
		self.conn = nil
		return
	}
	self.Log("Disconnected from database.")
	self.conn.Close()
}

func (self *tStore) Connect() error {
	self.Disconnect()
	cnx := ""
	if arg.Config.Http.DockerMode() {
		cnx = arg.Config.Db.ConnectString()
	} else {
		cnx = arg.Config.Db.ConnectStringLocalhost()
	}
	postgres, err := sql.Open("postgres", cnx)
	if err != nil {
		self.Error(fmt.Sprintf("Failed to connect to %s.\n\tError:%s", cnx, err.Error()))
		return err
	}
	err = postgres.Ping()
	if err != nil {
		self.Error(fmt.Sprintf("Ping failed on new postgres connection to %s. \n\tError:%s", cnx, err.Error()))
		return err
	} else {
		self.Log(fmt.Sprintf("Pinged postgres."))
	}
	self.conn = postgres
	self.Query = dbg.New(self.conn)
	self.Log("Succesfully connected to postgres.")
	return nil
}

func (self *tStore) PrintLogs() {
	logs := *self.GetLogs()
	for _, log := range logs {
		fmt.Println("db log:", log)
	}
	errs := *self.GetErrors()
	for _, err := range errs {
		fmt.Println("db err:", err)
	}
}
