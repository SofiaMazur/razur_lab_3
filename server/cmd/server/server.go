package main

import (
	"context"
	"net/http"
	"fmt"
	h "github.com/SofiaMazur/razur_lab_3/server/handlers"
)

type SiteServer struct {
	server *http.Server
	Senv *ServerEnv
	Handlers *h.Handlers
}

func (fs *SiteServer) Run() error {
	handlersCollection := map[string] http.HandlerFunc {
		"/users": fs.Handlers.HandleUsers,
		"/sites": fs.Handlers.HandleSites,
		"/user": fs.Handlers.GetUser,
		"/site": fs.Handlers.GetSite,
	}
	for route, handler := range handlersCollection {
		http.Handle(route, handler)
	}
	runnable := fs.Senv.Host + ":" + fmt.Sprint(fs.Senv.Port)
	fs.server = &http.Server{Addr: runnable}
	fmt.Printf("Server is running on port: %d, host: %s\n", fs.Senv.Port, fs.Senv.Host)
	return fs.server.ListenAndServe()
}

func (fs *SiteServer) Close() error {
	if fs.server == nil {
		return fmt.Errorf("Server was not started")
	}
	return fs.server.Shutdown(context.Background())
}
