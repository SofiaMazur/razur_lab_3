package uniqueStore

import (
	"database/sql"
	"github.com/SofiaMazur/razur_lab_3/server/tools"
	_ "github.com/lib/pq"
)

type UniqueStore struct {
	SStore *SiteStore
	UStore *UserStore
}

func NewUniqueStore(db *sql.DB) *UniqueStore {
	sstore := NewSiteStore(db)
	ustore := NewUserStore(db)
	return &UniqueStore{SStore: sstore, UStore: ustore}
}

func (gs *UniqueStore) ListSites() (*tools.Sites, error) {
	return gs.SStore.ListSites()
}

func (gs *UniqueStore) FindSiteByName(name string) (*tools.Sites, error) {
	return gs.SStore.FindSiteByName(name)
}

func (gs *UniqueStore) FindUserByName(name string) (*tools.Users, error) {
	return gs.UStore.FindUserByName(name)
}

func (gs *UniqueStore) CreateSite(name, topicKeyword string) error {
	return gs.SStore.CreateSite(name, topicKeyword)
}

func (gs *UniqueStore) ListUsers() (*tools.Users, error) {
	return gs.UStore.ListUsers()
}

func (gs *UniqueStore) CreateUser(username string, interests []string) error {
	return gs.UStore.CreateUser(username, interests)
}
