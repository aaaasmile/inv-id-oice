package app

import (
	"inv-id-oice/conf"
	"inv-id-oice/db"
	"inv-id-oice/web/app/models"
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
	status := http.StatusOK
	gh := models.NewModelHandler(conf.Current.Debug, ap.liteDB)
	if err := gh.HandleModel(w, req, &status); err != nil {
		log.Println("Error on process request: ", err)
		if status == http.StatusNotFound {
			http.Error(w, "404 - Not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
