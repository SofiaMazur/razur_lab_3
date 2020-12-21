package main

import (
	"github.com/SofiaMazur/razur_lab_3/server/uniqueStore"
	"github.com/SofiaMazur/razur_lab_3/server/handlers"
	"github.com/google/wire"
)


func NewServer(senv *ServerEnv) (*SiteServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	uniqueStoreUniqueStore := uniqueStore.NewUniqueStore(db)
	handlersHandlers := handlers.NewHandler(uniqueStoreUniqueStore)
	siteServer := &SiteServer{
		Handlers: handlersHandlers,
		Senv:     senv,
	}
	return siteServer, nil
}


var providers = wire.NewSet(uniqueStore.NewUniqueStore, handlers.NewHandler)
