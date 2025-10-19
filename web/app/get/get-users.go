package get

import (
	"html/template"
	"inv-id-oice/idl"
	"inv-id-oice/util"
	"log"
	"net/http"
)

func (gh *TypeGetHandler) handleUsers(w http.ResponseWriter) error {

	if gh.debug {
		log.Println("provides User View")
	}
	users := []idl.User{
		{Name: "John Doe", Email: "john.doe@example.com", Enabled: true},
		{Name: "Jane Smith", Email: "jane.smith@example.com", Enabled: true},
		{Name: "Robert Johnson", Email: "robert.johnson@example.com", Enabled: true},
	}
	pagectx := struct {
		Users []idl.User
	}{
		Users: users,
	}
	templName := "templates/app/views/users.html"

	tmpl := template.Must(template.New("Users").ParseFiles(util.GetFullPath(templName)))

	err := tmpl.ExecuteTemplate(w, "users", pagectx)

	return err
}
