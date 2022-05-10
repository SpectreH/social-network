package models

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	BirthDate string `json:"birthDate"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	AboutMe   string `json:"aboutMe"`
	Password  string `json:"password"`
}
