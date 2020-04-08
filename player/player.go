package player

import (
	"github.com/gorilla/websocket"
	"fmt"
)

type Message struct {
	Type int
	Content []byte
}

type Player struct {
	Id int
	conn *websocket.Conn
	Out chan Message
	In chan Message
}


func New(id int, conn *websocket.Conn) *Player {
	conn.SetCloseHandler(func (code int, message string) error {
		fmt.Printf("Player[%d] sent a [Close] message...", id)
		return nil
	})
	return &Player{
		Id: id, 
		conn: conn,
		Out: make(chan Message, 1024),
		In: make(chan Message, 1024),
	}
}

func (p *Player) ListenRead() {
	for {
		messageType, message, err := p.conn.ReadMessage()
		if err != nil {
			close(p.Out)
			break
		}
		p.Out <- Message{ messageType, message }
	}
}

func (p *Player) SendWrites() {
	for m := range p.In {
		err := p.conn.WriteMessage(m.Type, m.Content)
		if err != nil {
			fmt.Println(err)
		}
	}
	p.conn.Close()
}