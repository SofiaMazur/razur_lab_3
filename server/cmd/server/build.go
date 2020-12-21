package main

import (
	"github.com/google/wire"
	gs "github.com/SofiaMazur/razur_lab_3/server/uniqueStore"
	h "github.com/SofiaMazur/razur_lab_3/server/handlers"
)

var providers = wire.NewSet(gs.NewUniqueStore, h.NewHandler)

func NewServer(senv *ServerEnv) (*SiteServer, error) {
	wire.Build(
		NewDbConnection,
		providers,
		wire.Struct(new(SiteServer), "Handlers", "Senv"),
	)

	return nil, nil
}
