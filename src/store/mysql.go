package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/chenji-kael/go-bootstrap/src/api/datatype"
	"github.com/chenji-kael/go-bootstrap/src/helper"
)

const TimeFormat = "2006-01-02T15:04:05Z07:00"

type DbStore struct{
	db *sql.DB
}

func NewDbStore() *DbStore {
	s := DbStore{}
	db, err := sql.Open("mysql", helper.Config.UserDataSource)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %v", err))
	}
	helper.Logger.Infoln("Connected to database")

	_,err = db.Exec("CREATE DATABASE IF NOT EXISTS iam")
	if err != nil {
		panic(err)
	}

	_,err = db.Exec("USE iam")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table if not exists Ntp (ntpAddress varchar(40) primary key) default charset=utf8")
	if err != nil {
		panic(err)
	}

	db.Close()
	db, err = sql.Open("mysql", helper.Config.UserDataSource+"iam")
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %v", err))
	}
	s.db = db
	return &s
}

func (s *DbStore)InsertNtpRecord(ntpAddress string) error {
	_, err := s.db.Exec("delete from Ntp")
	if err != nil {
		helper.Logger.Errorln("clear ntp table failed", err.Error())
		return err
	}

	_, err = s.db.Exec("insert into Ntp values( ? )", ntpAddress)
	if err != nil {
		helper.Logger.Errorln("Error add ntp address", ntpAddress, err.Error())
		return err
	}
	return err
}

func (s *DbStore)GetNtpRecord() (NtpRecord, error) {
	var record NtpRecord
	err := s.db.QueryRow("select * from Ntp").Scan(&record.NtpAddress)
	return record, err
}

func (s *DbStore)Close() {
	s.db.Close()
}


