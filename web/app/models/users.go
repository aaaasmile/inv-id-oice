package models

import (
	"fmt"
	"html/template"
	"inv-id-oice/idl"
	"inv-id-oice/util"
	"log"
	"net/http"
)

var g_users = map[string]idl.User{
	"1": {Name: "John Doe", Email: "john.doe@example.com", Enabled: true, Id: 1},
	"2": {Name: "Jane Smith", Email: "jane.smith@example.com", Enabled: true, Id: 2},
	"3": {Name: "Robert Johnson", Email: "robert.johnson@example.com", Enabled: true, Id: 3},
}

func (mh *ModelHandler) handleUsers(w http.ResponseWriter) error {

	switch mh.req.Method {
	case "GET":
		id := mh.req.URL.Query().Get("id")
		if id != "" {
			return mh.viewOrEditSingleUser(id, w)
		} else {
			if mh.req.URL.Query().Get("new") == "1" {
				return mh.viewAddNewUser(w)
			}
			return mh.viewAllUsers(w)
		}
	case "PUT":
		id := mh.req.URL.Query().Get("id")
		if id != "" {
			return mh.updateSingleUser(id, w)
		} else {
			return fmt.Errorf("PUT: id %s to edit not found", id)
		}
	case "DELETE":
		id := mh.req.URL.Query().Get("id")
		if id != "" {
			return mh.deleteSingleUser(id, w)
		} else {
			return fmt.Errorf("DELETE id %s to delete not found", id)
		}
	case "POST":
		return mh.addUser(w)
	default:
		return fmt.Errorf("method %s not suported", mh.req.Method)
	}
}

func (mh *ModelHandler) viewAllUsers(w http.ResponseWriter) error {
	if mh.debug {
		log.Println("provides All Users View", mh.req.Method)
	}

	var users []idl.User
	for _, user := range g_users {
		users = append(users, user)
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

func (mh *ModelHandler) viewOrEditSingleUser(id string, w http.ResponseWriter) error {
	if mh.debug {
		log.Println("provides Single User View/Edit", mh.req.Method)
	}
	if _, ok := g_users[id]; !ok {
		return fmt.Errorf("key %s not found", id)
	}
	pagectx := g_users[id]
	templName := "templates/app/views/users.html"

	tmpl := template.Must(template.New("User").ParseFiles(util.GetFullPath(templName)))

	section_name := "single_user"
	edit := mh.req.URL.Query().Get("edit")
	if edit == "1" {
		log.Println("Editing user: ", id)
		section_name = "edit_user"
	} else {
		log.Println("View user: ", id)
	}

	err := tmpl.ExecuteTemplate(w, section_name, pagectx)

	return err
}

func (mh *ModelHandler) updateSingleUser(id string, w http.ResponseWriter) error {
	if mh.debug {
		log.Printf("Update user %s with id %s", mh.req.Method, id)
	}
	if v, ok := g_users[id]; ok {
		name := mh.req.PostFormValue("name")
		email := mh.req.PostFormValue("email")
		v.Name = name
		v.Email = email
		g_users[id] = v
	} else {
		return fmt.Errorf("key %s not found", id)
	}

	return mh.viewAllUsers(w)
}

func (mh *ModelHandler) deleteSingleUser(id string, w http.ResponseWriter) error {
	if mh.debug {
		log.Println("Single User Delete", mh.req.Method)
	}
	if _, ok := g_users[id]; !ok {
		return fmt.Errorf("key %s not found", id)
	}
	delete(g_users, id)

	return mh.viewAllUsers(w)
}

func (mh *ModelHandler) viewAddNewUser(w http.ResponseWriter) error {
	if mh.debug {
		log.Println("view add new user form", mh.req.Method)
	}

	pagectx := idl.User{}
	templName := "templates/app/views/users.html"

	tmpl := template.Must(template.New("User").ParseFiles(util.GetFullPath(templName)))

	section_name := "new_user"

	return tmpl.ExecuteTemplate(w, section_name, pagectx)
}

func (mh *ModelHandler) addUser(w http.ResponseWriter) error {
	if mh.debug {
		log.Println("New User", mh.req.Method)
	}
	nextId := 0
	for key := range g_users {
		uu := g_users[key]
		if uu.Id > nextId {
			nextId = uu.Id
		}
	}
	nextId += 1

	v := idl.User{}
	name := mh.req.PostFormValue("name")
	email := mh.req.PostFormValue("email")
	v.Name = name
	v.Email = email
	// TODO Error handling on validation

	v.Id = nextId
	id := fmt.Sprintf("%d", nextId)
	g_users[id] = v

	return mh.viewAllUsers(w)
}
