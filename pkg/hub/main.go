package hub

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type hub struct {
	connections map[string]*websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewHub() *hub {
	return &hub{
		connections: make(map[string]*websocket.Conn),
	}
}

func (h *hub) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	query := r.URL.Query()
	if !query.Has("planetId") || !query.Has("password") {
		log.Println("planetId or password missing")
		cm := websocket.FormatCloseMessage(400, "planetId or password missing")
		conn.WriteMessage(websocket.CloseMessage, cm)
		conn.Close()
		return
	}

	planetId := query.Get("planetId")
	h.connections[planetId] = conn

	log.Println("connected: ", planetId)

	defer conn.Close()
	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			log.Println("read: ", err)
			break
		}

		log.Printf("%s sent: %s\n", planetId, string(msg))
	}

}
