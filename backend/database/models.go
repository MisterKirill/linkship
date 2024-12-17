package database

type AuthenticationUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReadUser struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Bio         string `json:"bio"`
	Color       string `json:"color"`
}

type User struct {
	ReadUser
	Id int `json:"id"`
}

type CreateLink struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Color string `json:"color"`
}

type ReadLink struct {
	CreateLink
	Id int `json:"id"`
}

type Link struct {
	ReadLink
	UserId int `json:"user_id"`
}
