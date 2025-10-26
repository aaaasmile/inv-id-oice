package models

import (
	"fmt"
	"html/template"
	"inv-id-oice/util"
	"log"
	"net/http"
)

func (mh *ModelHandler) handleSampleTable(w http.ResponseWriter) error {

	switch mh.req.Method {
	case "GET":
		return mh.viewSampleTable(w)
	default:
		return fmt.Errorf("method %s not suported", mh.req.Method)
	}
}

func (mh *ModelHandler) viewSampleTable(w http.ResponseWriter) error {
	if mh.debug {
		log.Println("show SampleTable", mh.req.Method)
	}

	pagectx := struct {
	}{}
	templName := "templates/app/views/sample_table.html"

	tmpl := template.Must(template.New("SampleTable").ParseFiles(util.GetFullPath(templName)))

	err := tmpl.ExecuteTemplate(w, "sample_table", pagectx)

	return err
}
