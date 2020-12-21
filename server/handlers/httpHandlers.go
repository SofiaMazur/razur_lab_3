package handlers

import (
	"net/http"
	"encoding/json"
	
	gs "github.com/SofiaMazur/razur_lab_3/server/uniqueStore"
	"github.com/SofiaMazur/razur_lab_3/server/tools"
)

type Handlers struct {
	db *gs.UniqueStore
}

func (h *Handlers) HandleUsers(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getUsers(h.db, rw, req)
	} else if req.Method == "POST" {
		addUser(h.db, rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) HandleSites(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getSites(h.db, rw, req)
	} else if req.Method == "POST" {
		addSite(h.db, rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

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

func (h *Handlers) GetSite(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var resName tools.ResponseName
	if err := json.NewDecoder(req.Body).Decode(&resName); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
		return
	}
	if res, err := h.db.FindSiteByName(resName.Name); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}
