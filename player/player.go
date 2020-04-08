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
	p := &Player{
		Id: id, 
		conn: conn,
		Out: make(chan Message, 1024),
		In: make(chan Message, 1024),
	}
	p.conn.SetCloseHandler(func (code int, message string) error {
		//close(p.Out)
		return nil
	})
	return p
}

func (p *Player) ListenRead() {
	for {
		messageType, message, err := p.conn.ReadMessage()
		if err != nil {
			fmt.Printf("Player[%d] sent a [Close] message...\n", p.Id)
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