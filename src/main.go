package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	. "github.com/chenji-kael/go-bootstrap/src/api"
	"github.com/chenji-kael/go-bootstrap/src/helper"
	"github.com/chenji-kael/go-bootstrap/src/store"
	"github.com/chenji-kael/go-bootstrap/src/mideleware/cors"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
	"strings"
)

var AllRoutes []map[string]string

func RegisterRequests(r *mux.Router) {
	r.HandleFunc("/", ListAPI).Methods("GET")
	r.HandleFunc("/api", ListAPI).Methods("GET")
	r.HandleFunc("/api/v1/manage/ntp", SetNTP).Methods("POST")
	r.HandleFunc("/api/v1/manage/ntp", GetNTP).Methods("GET")
}

func ListAPI(w http.ResponseWriter, r *http.Request) {
	rtjson, _ := json.Marshal(AllRoutes)
	fmt.Fprintf(w, string(rtjson))
	return
}

func ShowAPI(r *mux.Router) {
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			fmt.Println(err)
			return err
		}
		m, err := route.GetMethods()
		if err != nil {
			return err
		}
		sm := strings.Join(m, ",")
		SRoute := map[string]string{"method": sm, "path": t}
		AllRoutes = append(AllRoutes, SRoute)
		return nil
	})

}

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedHeaders:     []string{"*"},
		OptionsPassthrough: false,
		AllowCredentials:   true,
	})
	helper.SetupConfig()
	r := mux.NewRouter()
	r.StrictSlash(true)
	n := negroni.New() // Includes some default middlewares
	// add log middleware
	logrusMiddleWare := negronilogrus.NewMiddleware()
	file, err := helper.OpenAccessLogFile()
	if err == nil {
		logrusMiddleWare.Logger.Out = file
		n.Use(logrusMiddleWare)
		defer file.Close()
	}
	RegisterRequests(r)
	ShowAPI(r)
	store.Db = store.NewDbStore()
	defer store.Db.Close()
	n.Use(c)
	recovery := negroni.NewRecovery()
	recovery.Formatter = &negroni.HTMLPanicFormatter{}
	n.Use(recovery)
	n.UseHandler(r)
	http.ListenAndServe(":8080", n)

}
