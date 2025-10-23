package models

import (
	"fmt"
	"html/template"
	"inv-id-oice/util"
	"log"
	"net/http"
)

func (mh *ModelHandler) handleLogin(w http.ResponseWriter) error {

	switch mh.req.Method {
	case "GET":
		return mh.viewLogin(w)
	default:
		return fmt.Errorf("method %s not suported", mh.req.Method)
	}
}

func (mh *ModelHandler) viewLogin(w http.ResponseWriter) error {
	if mh.debug {
		log.Println("show login", mh.req.Method)
	}

	pagectx := struct {
	}{}
	templName := "templates/app/views/login.html"

	tmpl := template.Must(template.New("Login").ParseFiles(util.GetFullPath(templName)))

	err := tmpl.ExecuteTemplate(w, "login", pagectx)

	return err
}
