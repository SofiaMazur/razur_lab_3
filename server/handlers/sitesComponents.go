package handlers

import (
	"encoding/json"
	"net/http"
	
	"github.com/SofiaMazur/razur_lab_3/server/tools"
	gs "github.com/SofiaMazur/razur_lab_3/server/uniqueStore"
)

func addForum(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	var forum tools.Forum
	if err := json.NewDecoder(req.Body).Decode(&forum); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
		return
	}
	if err := db.CreateForum(forum.Name, forum.Topic); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		rw.WriteHeader(http.StatusCreated)
	}
}

func getForums(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	if res, err := db.ListForums(); err != nil {
		tools.WriteJsonInternalError(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}
