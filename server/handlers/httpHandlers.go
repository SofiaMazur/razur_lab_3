package handlers

import (
	"net/http"
	"encoding/json"
	
	gs "github.com/SofiaMazur/razur_lab_3/server/uniqueStore"
	"github.com/SofiaMazur/razur_lab_3/server/tools"
)

// Handlers for server routes
type Handlers struct {
	db *gs.UniqueStore
}

// HandleUsers for POST and GET methods
func (h *Handlers) HandleUsers(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getUsers(h.db, rw, req)
	} else if req.Method == "POST" {
		addUser(h.db, rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// HandleForums for POST and GET methods
func (h *Handlers) HandleForums(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getForums(h.db, rw, req)
	} else if req.Method == "POST" {
		addForum(h.db, rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// GetUser for POST method
func (h *Handlers) GetUser(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var resName tools.ResponseName
	if err := json.NewDecoder(req.Body).Decode(&resName); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
		return
	}
	if res, err := h.db.FindUserByName(resName.Name); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}

// GetForum for POST method
func (h *Handlers) GetForum(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var resName tools.ResponseName
	if err := json.NewDecoder(req.Body).Decode(&resName); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
		return
	}
	if res, err := h.db.FindForumByName(resName.Name); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}
