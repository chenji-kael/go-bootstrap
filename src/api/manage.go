package api

import (
	"encoding/json"
	. "github.com/chenji-kael/go-bootstrap/src/api/datatype"
	"github.com/chenji-kael/go-bootstrap/src/store"
	. "github.com/chenji-kael/go-bootstrap/src/error"
	"github.com/chenji-kael/go-bootstrap/src/helper"
	"database/sql"
	"io/ioutil"
	"net/http"
)

func SetNTP(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	req := &QueryRequestNtp{}
	err := json.Unmarshal(body, req)
	if err != nil {
		WriteErrorResponse(w, r, ErrJsonDecodeFailed)
		return
	}

	err = store.Db.InsertNtpRecord(req.NtpAddress)
	if err != nil {
		helper.Logger.Infoln("failed SetNTP for query:", req)
		WriteErrorResponse(w, r, err)
		return
	}
	WriteSuccessResponse(w, nil)
}

func GetNTP(w http.ResponseWriter, r *http.Request) {
	//param1 := r.URL.Query().Get("param1")
	//if param1 != "" {
	//	// ... process it, will be the first (only) if multiple were given
	//	// note: if they pass in like ?param1=&param2= param1 will also be "" :|
	//}
	record, err := store.Db.GetNtpRecord()
	if err != nil && err != sql.ErrNoRows {
		helper.Logger.Infoln("failed GetNTP for query:", r)
		WriteErrorResponse(w, r, err)
		return
	}
	WriteSuccessResponse(w, EncodeResponse(record))
}


