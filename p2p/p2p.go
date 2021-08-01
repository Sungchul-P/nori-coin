package p2p

import (
	"github.com/Sungchul-P/nori-coin/utils"
	"github.com/gorilla/websocket"
	"net/http"
)

var conns []*websocket.Conn
var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	conns = append(conns, conn)
	utils.HandleErr(err)
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			break
		}
		for _, aConn := range conns {
			if aConn != conn {
				utils.HandleErr(aConn.WriteMessage(websocket.TextMessage, p))
			}
		}
	}
}
