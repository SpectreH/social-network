package socket

import (
	"database/sql"
	"encoding/json"
	"log"
	"social-network/internal/config"
	"social-network/internal/database"
	"social-network/internal/database/sqlite"
	"social-network/internal/models"
	"strconv"
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

	toId, _ := strconv.Atoi(socketMessage.Dest)

	avatar, err := sr.db.GetUserAvatar(sr.id)
	if err != nil {
		panic(err)
	}

	socketMessage.Message = config.FOLLOW_REQUEST_MESSAGE
	socketMessage.Avatar = config.AVATAR_PATH_URL + avatar
	socketMessage.To = toId
	socketMessage.Source = sr.id
	socketMessage.SourceName = sr.name
	socketMessage.Date = time.Now()

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
