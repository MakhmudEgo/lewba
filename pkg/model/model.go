package model

type PairToken struct {
	Access     string
	Refresh    string
	RefreshExp int64
}

type User struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
