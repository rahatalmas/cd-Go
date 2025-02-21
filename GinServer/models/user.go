package models

type Auth struct {
	Name     string `json:"user_name"`
	Password string `json:"password"`
}

type User struct {
	Id           int    `json:"user_id"`
	Name         string `json:"user_name"`
	Password     string `json:"user_password"`
	Picture      string `json:"user_picture,omitempty"`
	Contact      string `json:"user_contact"`
	Role         string `json:"role_key"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
