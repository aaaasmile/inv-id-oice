package app

import (
	"fmt"
	"inv-id-oice/conf"
	"inv-id-oice/db"
	"inv-id-oice/web/app/get"
	"inv-id-oice/web/app/post"
	"log"
	"net/http"
)

type App struct {
	liteDB *db.LiteDB
}

func NewApp() (*App, error) {
	res := &App{}
	var err error
	if res.liteDB, err = db.OpenSqliteDatabase(conf.Current.Database.DbFileName,
		conf.Current.Database.SQLDebug); err != nil {
		return nil, err
	}
	return res, nil
}

func (ap *App) APiHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		status := http.StatusOK
		gh := get.NewTypeGetHandler(conf.Current.Debug, ap.liteDB)
		if err := gh.HandleGet(w, req, &status); err != nil {
			log.Println("Error on process request: ", err)
			if status == http.StatusNotFound {
				http.Error(w, "404 - Not found", http.StatusNotFound)
			} else {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}
	case "POST":
		ph := post.NewPostHandler(conf.Current.Debug, ap.liteDB)
		if err := ph.HandlePost(w, req); err != nil {
			log.Println("[POST] Error: ", err)
			http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
			return
		}
	}
}
