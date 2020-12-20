package generalStore

import (
	"database/sql"
	"github.com/SofiaMazur/razur_lab_3/server/tools"
	_ "github.com/lib/pq"
)

type UniqueStore struct {
	FStore *ForumStore
	UStore *UserStore
}

func NewGeneralStore(db *sql.DB) *GeneralStore {
	fstore := NewForumStore(db)
	ustore := NewUserStore(db)
	return &UniqueStore{FStore: fstore, UStore: ustore}
}

func (gs *UniqueStore) ListForums() (*tools.Forums, error) {
	return gs.FStore.ListForums()
}

func (gs *UniqueStore) FindForumByName(name string) (*tools.Forums, error) {
	return gs.FStore.FindForumByName(name)
}

func (gs *UniqueStore) FindUserByName(name string) (*tools.Users, error) {
	return gs.UStore.FindUserByName(name)
}

func (gs *UniqueStore) CreateForum(name, topicKeyword string) error {
	return gs.FStore.CreateForum(name, topicKeyword)
}

func (gs *UniqueStore) ListUsers() (*tools.Users, error) {
	return gs.UStore.ListUsers()
}

func (gs *UniqueStore) CreateUser(username string, interests []string) error {
	return gs.UStore.CreateUser(username, interests)
}
