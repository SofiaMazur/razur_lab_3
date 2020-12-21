package handlers

import (
	"encoding/json"
	"net/http"
	
	"github.com/SofiaMazur/razur_lab_3/server/tools"
	gs "github.com/SofiaMazur/razur_lab_3/server/uniqueStore"
)

func addUser(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	var user tools.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
		return
	}
	if err := db.CreateUser(user.Name, user.Interests); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		rw.WriteHeader(http.StatusCreated)
	}
}

func getUsers(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	if res, err := db.ListUsers(); err != nil {
		tools.WriteJsonInternalError(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}
