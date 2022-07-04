package socket

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"social-network/internal/config"
	"social-network/internal/database"
	"social-network/internal/database/sqlite"
	"social-network/internal/models"
	"time"

	"github.com/gorilla/websocket"
)

// SocketReader is the socket reader model
type SocketReader struct {
	db   database.DatabaseRepo
	conn *websocket.Conn
	name string
	id   int
}

var savedSocketReader []*SocketReader

// SetSocketReader creates a new socket reader
func SetSocketReader(conn *sql.DB) *SocketReader {
	return &SocketReader{
		db: sqlite.SetSqliteRepo(conn),
	}
}

// AppendNewConnection appends a new user connection into array
func (sr *SocketReader) AppendNewConnection(conn *websocket.Conn, name string, id int) {
	if savedSocketReader == nil {
		savedSocketReader = make([]*SocketReader, 0)
	}

	ptrSocketReader := &SocketReader{
		conn: conn,
		db:   sr.db,
		name: name,
		id:   id,
	}

	savedSocketReader = append(savedSocketReader, ptrSocketReader)

	ptrSocketReader.startThread()
}

func (sr *SocketReader) removeMultipleConnection() {
	for _, socket := range savedSocketReader {
		if socket == sr {
			continue
		}

		if socket.id == sr.id {
			socket.closeConnection()
		}
	}
}

// startThread connects user to chat
func (sr *SocketReader) startThread() {
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				log.Println(err)
			}
			log.Println("thread socketreader finish")
			sr.closeConnection()
		}()

		sr.removeMultipleConnection()

		for {
			sr.read()
		}
	}()
}

// broadcastMultiple sends message to multiple users from array
func (sr *SocketReader) broadcastMultiple(message models.SocketMessage, ids []int) {
	for _, g := range savedSocketReader {
		for _, i := range ids {
			if sr.id == i {
				continue
			}

			if i == g.id {
				g.writeMsg(message)
			}
		}
	}
}

// broadcast sends message to user
func (sr *SocketReader) broadcast(message models.SocketMessage) {
	for _, g := range savedSocketReader {
		if message.To == g.id {
			g.writeMsg(message)
		}
	}
}

// read reads messages from user
func (sr *SocketReader) read() {
	_, b, err := sr.conn.ReadMessage()
	if err != nil {
		panic(err)
	}

	if string(b) == "heart_beat" {
		return
	}

	socketMessage := models.SocketMessage{}

	err = json.Unmarshal(b, &socketMessage)
	if err != nil {
		panic(err)
	}

	if socketMessage.Type == config.NEW_EVENT_REQUEST_TYPE {
		parts, err := sr.db.GetGroupParticipants(socketMessage.GroupId)
		if err != nil {
			panic(err)
		}

		group, err := sr.db.GetGroup(socketMessage.GroupId)
		if err != nil {
			panic(err)
		}

		socketMessage.Message = fmt.Sprintf("Has new event %s what will be held on %s", socketMessage.EventName, socketMessage.EventDate)
		socketMessage.Avatar = group.Picture
		socketMessage.SourceName = group.Title
		socketMessage.Date = time.Now()

		sr.broadcastMultiple(socketMessage, parts)
		return
	}

	avatar, err := sr.db.GetUserAvatar(sr.id)
	if err != nil {
		panic(err)
	}
	socketMessage.Avatar = config.AVATAR_PATH_URL + avatar

	if socketMessage.Type == config.GROUP_FOLLOW_REQUEST_TYPE {
		socketMessage.Message = config.GROUP_FOLLOW_REQUEST_MESSAGE + socketMessage.GroupName
	} else if socketMessage.Type == config.FOLLOW_REQUEST_TYPE {
		socketMessage.Message = config.FOLLOW_REQUEST_MESSAGE
	} else if socketMessage.Type == config.GROUP_INVITE_TYPE {
		socketMessage.Message = config.GROUP_INVITE_MESSAGE + socketMessage.GroupName
	}

	socketMessage.Source = sr.id
	socketMessage.SourceName = sr.name
	socketMessage.Date = time.Now()

	if socketMessage.IsGroupChat && socketMessage.Type == config.NEW_MESSAGE_TYPE {
		parts, err := sr.db.GetGroupParticipantsByChat(socketMessage.ChatId)
		if err != nil {
			panic(err)
		}

		sr.broadcastMultiple(socketMessage, parts)
		return
	}

	sr.broadcast(socketMessage)
}

// writeMsg writes message to user
func (sr *SocketReader) writeMsg(message models.SocketMessage) {
	sr.conn.WriteJSON(message)
}

// closeConnection disconnects user and removes from the array
func (sr *SocketReader) closeConnection() {
	for i, socket := range savedSocketReader {
		if socket == sr {
			savedSocketReader = append(savedSocketReader[:i], savedSocketReader[i+1:]...)
			sr.conn.Close()
			return
		}
	}
}
