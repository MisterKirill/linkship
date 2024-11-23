package database

type ReadUser struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Bio         string `json:"bio"`
}

type User struct {
	ReadUser
	Id int `json:"id"`
}

type AuthenticationUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateLink struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ReadLink struct {
	CreateLink
	Id int `json:"id"`
}

type Link struct {
	ReadLink
	UserId int `json:"user_id"`
}
