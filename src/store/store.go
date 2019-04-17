package store

import (
	_ "github.com/go-sql-driver/mysql"
	. "github.com/chenji-kael/go-bootstrap/src/api/datatype"
)

var Db backendstore

type backendstore interface {
	InsertNtpRecord(ntpAddress string) error
	GetNtpRecord() (NtpRecord, error)
	Close()
}