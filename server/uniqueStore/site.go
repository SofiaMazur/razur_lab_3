package uniqueStore

import (
	"database/sql"
	"fmt"
	"github.com/SofiaMazur/razur_lab_3/server/tools"
	_ "github.com/lib/pq"
)

type SiteStore struct {
	Db *sql.DB
}

func NewSiteStore(db *sql.DB) *SiteStore {
	return &SiteStore{Db: db}
}

func (s *SiteStore) ListSites() (*tools.Sites, error) {
	rows, err := s.Db.Query("SELECT * FROM sites")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*tools.Site
	for rows.Next() {
		var f tools.Site
		if err := rows.Scan(&f.Id, &f.Name, &f.Topic); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}

	var fullSites []*tools.Site
	if res == nil {
		fullSites = make([]*tools.Site, 0)
	} else {
		for i := 0; i < len(res); i++ {
			users, err := s.GetSiteUsersByID(res[i].Id)
			if err != nil {
				return nil, err
			}
			fullSite := tools.Site{
				Id:    res[i].Id,
				Name:  res[i].Name,
				Topic: res[i].Topic,
				Users: users}
			fullSites = append(fullSites, &fullSite)
		}
	}

	result := &tools.Sites{fullSites}
	return result, err
}

func (s *SiteStore) FindSiteByName(name string) (*tools.Sites, error) {
	var textError string
	var err error
	var fullSites []*tools.Site

	if len(name) == 0 {
		textError = "Site name is not provided"
		err = fmt.Errorf(textError)
		fullSites = make([]*tools.Site, 0)
		return nil, err
	}
	rows, err := s.Db.Query(`SELECT * FROM sites where name = $1`, name)
	if err != nil {
		textError = "There is no such site"
		err = fmt.Errorf(textError)
		return nil, err
	}

	defer rows.Close()

	var res []*tools.Site
	for rows.Next() {
		var f tools.Site
		if err = rows.Scan(&f.Id, &f.Name, &f.Topic); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}
	if res == nil {
		textError = "No such site"
		err = fmt.Errorf(textError)
		return nil, err
	}
	for i := 0; i < len(res); i++ {
		users, err := s.GetSiteUsersByID(res[i].Id)
		if err != nil {
			return nil, err
		}
		fullSite := tools.Site{
			Id:    res[i].Id,
			Name:  res[i].Name,
			Topic: res[i].Topic,
			Users: users}
		fullSites = append(fullSites, &fullSite)
	}

	result := &tools.Sites{fullSites}
	return result, nil
}

func (s *SiteStore) FindSiteByTopic(name string) ([]*tools.Site, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("Topic name is not provided")
	}
	rows, err := s.Db.Query(`SELECT * FROM sites where topicKeyword = $1`, name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*tools.Site
	for rows.Next() {
		var f tools.Site
		if err := rows.Scan(&f.Id, &f.Name, &f.Topic); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}
	if res == nil {
		res = make([]*tools.Site, 0)
		return res, fmt.Errorf("no such site")
	}
	return res, nil
}

func (s *SiteStore) CreateSite(name, topicKeyword string) error {
	if len(name) == 0 {
		return fmt.Errorf("Site name is not provided")
	}
	if len(topicKeyword) == 0 {
		return fmt.Errorf("Topic keyword name is not provided")
	}
	_, err := s.Db.Exec(`INSERT INTO sites (name, topicKeyword) VALUES ($1, $2)`, name, topicKeyword)
	if err != nil {
		return fmt.Errorf("Site with this name or topic already exists")
	}
	sites, err := s.FindSiteByName(name)
	_, err = s.Db.Exec(`INSERT INTO usersList (sitesID) VALUES ($1)`, sites.SitesArr[0].Id)
	return err
}

func (s *SiteStore) GetSiteUsersByID(id int) ([]string, error) {
	if id < 1 {
		return nil, fmt.Errorf("ID is incorrect")
	}
	rows, err := s.Db.Query(`
	select
		users.name
	from
		sites
	left join
		usersList
	on
		usersList.sitesID = sites.id
	left join
		users
	on
		users.id = usersList.userID
	where
		sites.id = $1
	GROUP BY
		users.id
	HAVING users.id is not NULL
	`,
		id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []string
	for rows.Next() {
		var u string
		if err := rows.Scan(&u); err != nil {
			return nil, err
		}
		if u != "" {
			res = append(res, u)
		}
	}
	if res == nil {
		res = make([]string, 0)
	}

	return res, nil
}

func (s *SiteStore) AddUserToSite(idSite, idUser int) error {
	_, err := s.Db.Exec(`INSERT INTO usersList (sitesID, userID) VALUES ($1, $2)`, idSite, idUser)
	return err
}
