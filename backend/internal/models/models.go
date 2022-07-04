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

type ChatMessage struct {
	Id             int       `json:"id"`
	ProfilePicture string    `json:"profilePicture"`
	ChatId         int       `json:"chatId"`
	AuthorId       int       `json:"author"`
	AuthorName     string    `json:"authorName"`
	Content        string    `json:"text"`
	CreatedAt      time.Time `json:"time"`
}

type Chat struct {
	Id        int           `json:"id"`
	GroupChat bool          `json:"isGroupChat"`
	DestId    int           `json:"destId"`
	Name      string        `json:"name"`
	Avatar    string        `json:"avatar"`
	Messages  []ChatMessage `json:"messages"`
}

type Group struct {
	Id             int          `json:"id"`
	ChatId         int          `json:"chatId"`
	CreatorId      int          `json:"creatorId"`
	Picture        string       `json:"picture"`
	Title          string       `json:"title"`
	Description    string       `json:"description"`
	Private        bool         `json:"private"`
	CreatedAt      time.Time    `json:"createdAt"`
	TotalFollowers int          `json:"totalFollowers"`
	TotalPosts     int          `json:"totalPosts"`
	Following      bool         `json:"isFollowing"`
	Invite         bool         `json:"invite"`
	IsMyGroup      bool         `json:"isMyGroup"`
	Events         []Event      `json:"events"`
	Posts          []PostInside `json:"posts"`
}

type Event struct {
	Id           int                `json:"id"`
	GroupId      int                `json:"groupId"`
	Picture      string             `json:"picture"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	Date         time.Time          `json:"date"`
	CreatedAt    time.Time          `json:"createdAt"`
	Participants []EventParticipant `json:"participants"`
}

type EventParticipant struct {
	Id            int  `json:"id"`
	ParticipantId int  `json:"participantId"`
	EventId       int  `json:"eventId"`
	WillAttend    bool `json:"willAttend"`
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
	Avatar      string    `json:"avatar"`
	To          int       `json:"dest"`
	EventName   string    `json:"eventName"`
	EventDate   string    `json:"eventDate"`
	EventId     int       `json:"eventId"`
	ChatId      int       `json:"chatId"`
	GroupName   string    `json:"groupName"`
	GroupId     int       `json:"groupId"`
	IsGroupChat bool      `json:"isGroupChat"`
	Source      int       `json:"authorId"`
	SourceName  string    `json:"author"`
	Message     string    `json:"sub"`
	Type        string    `json:"type"`
	Date        time.Time `json:"date"`
}
