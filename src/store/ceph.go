package store

//import (
//	. "github.com/chenji-kael/go-bootstrap/src/api/datatype"
//	. "github.com/chenji-kael/go-bootstrap/src/error"
//	"github.com/chenji-kael/go-bootstrap/src/helper"
//	"time"
//	"github.com/journeymidnight/go-ceph/rados"
//	"encoding/json"
//	"math"
//	"fmt"
//	"strings"
//	"sort"
//	"strconv"
//	"errors"
//)
//
//const (
//	NtpTable = "DEADBEAF_NTP_TABLE"
//)
//
//type CephStore struct{
//	conn *rados.Conn
//	ioctx *rados.IOContext
//}
//
//func NewCephStore() *CephStore {
//	conn, err := rados.NewConn()
//	conn.ReadDefaultConfigFile()
//	conn.Connect()
//	a := CephStore{}
//	a.conn = conn
//	pools, err := conn.ListPools()
//	if err != nil {
//		helper.Logger.Println("list pools error:", err)
//		panic(err)
//	}
//	helper.Logger.Warnln("list pools result:", pools)
//
//	for _, v := range pools {
//		if v == "nier" {
//			a.ioctx, err = conn.OpenIOContext("nier")
//			if err != nil {
//				helper.Logger.Println("open pool nier error:", err)
//				panic(err)
//			}
//			a.TouchObjectIfNotExist(NtpTable)
//			return &a
//		}
//		helper.Logger.Println("list pool:", v)
//	}
//	err = conn.MakePool("nier")
//	if err != nil {
//		helper.Logger.Println("make pool nier:", err)
//		panic(err)
//	}
//	a.ioctx, err = conn.OpenIOContext("nier")
//	if err != nil {
//		helper.Logger.Println("open pool nier error:", err)
//		panic(err)
//	}
//	a.TouchObjectIfNotExist(NtpTable)
//
//
//	return &a
//}
//
//func (s *CephStore)TouchObjectIfNotExist(name string) {
//	_, err := s.ioctx.Stat(name)
//	if err != nil && err.Error() == "rados: No such file or directory" {
//		err = s.ioctx.Write(name, []byte("let`s rock"), 0)
//		if err != nil {
//			panic(fmt.Sprintf("touch object %s failed", name))
//		}
//		err = s.ioctx.Truncate(name, 0)
//		if err != nil {
//			panic(fmt.Sprintf("truncate object %s failed", name))
//		}
//	}
//}
//
//func (s *CephStore)InsertNtpRecord(ntpAddress string) error {
//	var record NtpRecord
//	record.NtpAddress = ntpAddress
//	data, _ := json.Marshal(record)
//	err := s.ioctx.SetOmap(NtpTable, map[string][]byte{
//		"ntp":data,
//	})
//
//	if err != nil {
//		helper.Logger.Errorln("Error add ntp address", ntpAddress, err.Error())
//		return err
//	}
//	return err
//}
//
//func (s *CephStore)GetNtpRecord() (NtpRecord, error) {
//	var record NtpRecord
//	results, err := s.ioctx.GetOmapValuesByKeys(NtpTable,[]string{"ntp"})
//	if err != nil {
//		return record, err
//	}
//	if len(results) > 0 {
//		err = json.Unmarshal(results["ntp"], &record)
//		return record, err
//	}
//	return record, nil
//}
//
//func (s *CephStore)Close() {
//	s.conn.Shutdown()
//}