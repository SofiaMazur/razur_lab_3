package tools

type User struct {
	Id        int      `json:"-"`
	Name string `json:"name"`
	Interests []string `json:"interests"`
}

type Site struct {
	Id      int    `json:"-"`
	Name string `json:"name"`
	Topic string `json: "topic"`
	Users []string `json:"users"`
}

type Sites struct {
	ForumsArr []*Forum `json:"forums"`
}

type ResponseName struct {
	Name string `json:"name"`
}

type Users struct {
	UsersArr []*User `json:"users"`
}

type errorObject struct {
	Message string `json:"message"`
}
