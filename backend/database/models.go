package database

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Bio         string `json:"bio"`
}

type AuthenticationUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
