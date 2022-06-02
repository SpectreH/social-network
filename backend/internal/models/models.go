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
	Id             int          `json:"id"`
	FirstName      string       `json:"firstName"`
	LastName       string       `json:"lastName"`
	Email          string       `json:"email"`
	BirthDate      string       `json:"birthDate"`
	Nickname       string       `json:"nickname"`
	Avatar         string       `json:"avatar"`
	AboutMe        string       `json:"aboutMe"`
	Private        bool         `json:"private"`
	Following      bool         `json:"following"`
	IsMyProfile    bool         `json:"isMyProfile"`
	TotalFollowers int          `json:"totalFollowers"`
	TotalFollows   int          `json:"totalFollows"`
	TotalPosts     int          `json:"totalPosts"`
	Followers      []Follow     `json:"followers"`
	Posts          []PostInside `json:"posts"`
}

type Follow struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Type      string `json:"type"`
	Followers int    `json:"totalFollowers"`
	Avatar    string `json:"avatar"`
	Selected  bool   `json:"selected"`
}

type PostInside struct {
	Post     Post      `json:"post"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Id     string `json:"id"`
	AuthId int    `json:"authorId"`
	PostId int    `json:"postId"`
	Author
	Picture   string    `json:"picture"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type Author struct {
	Avatar    string `json:"avatar"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Post struct {
	Id     int `json:"id"`
	AuthId int `json:"authorId"`
	Author
	GroupId     int    `json:"groupId"`
	GroupTitle  string `json:"groupTitle"`
	GroupAvatar string `json:"groupAvatar"`

	ShareId    int       `json:"shareId"`
	Picture    string    `json:"picture"`
	Paragraphs []string  `json:"paragraphs"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
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
