package handlers

import (
	gs "github.com/SofiaMazur/razur_lab_3/server/uniqueStore"
)

func NewHandler(gs *gs.UniqueStore) *Handlers {
	return &Handlers{gs}
}
