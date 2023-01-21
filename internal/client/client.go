package client

type User struct {
	Username string `json:"login"`
	Password string `json:"password"`
	Passhash string
}
