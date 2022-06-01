package models

import "time"

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	BirthDate string `json:"birthDate"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	AboutMe   string `json:"aboutMe"`
	Password  string `json:"password"`
	Private   bool   `json:"private"`
}

type UserProfile struct {
	Id             int      `json:"id"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Email          string   `json:"email"`
	BirthDate      string   `json:"birthDate"`
	Nickname       string   `json:"nickname"`
	Avatar         string   `json:"avatar"`
	AboutMe        string   `json:"aboutMe"`
	Private        bool     `json:"private"`
	Following      bool     `json:"following"`
	IsMyProfile    bool     `json:"isMyProfile"`
	TotalFollowers int      `json:"totalFollowers"`
	TotalFollows   int      `json:"totalFollows"`
	TotalPosts     int      `json:"totalPosts"`
	Followers      []Follow `json:"followers"`
}

type Follow struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Type      string `json:"type"`
	Followers int    `json:"totalFollowers"`
	Avatar    string `json:"avatar"`
}

type Post struct {
	Id        string
	AuthId    int
	ShareId   int
	Title     string
	Content   []string
	CreatedAt time.Time
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FormValidationResponse struct {
	OK      bool   `json:"ok"`
	Input   string `json:"input"`
	Message string `json:"message"`
	Token   string `json:"token"`
	Data    string `json:"data"`
}

type SocketMessage struct {
	Dest string `json:"dest"`

	Avatar string `json:"avatar"`

	To int `json:"to"`

	Source     int    `json:"authorId"`
	SourceName string `json:"author"`
	Message    string `json:"sub"`

	Type string    `json:"type"`
	Date time.Time `json:"date"`
}
