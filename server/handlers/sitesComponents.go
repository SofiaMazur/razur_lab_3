package handlers

import (
	"encoding/json"
	"net/http"
	
	"github.com/SofiaMazur/razur_lab_3/server/tools"
	gs "github.com/SofiaMazur/razur_lab_3/server/uniqueStore"
)

func addSite(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	var site tools.Site
	if err := json.NewDecoder(req.Body).Decode(&site); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
		return
	}
	if err := db.CreateSite(site.Name, site.Topic); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		rw.WriteHeader(http.StatusCreated)
	}
}

func getSites(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	if res, err := db.ListSites(); err != nil {
		tools.WriteJsonInternalError(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}
