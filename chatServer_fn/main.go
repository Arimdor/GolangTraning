package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}
type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: conn.id + " se ha conectado."})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: conn.id + " se ha desconecto"})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}
func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}
func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()
	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func main() {
	fmt.Println("Starting application on port 80...")
	go manager.start()
	mux := mux.NewRouter()
	mux.HandleFunc("/ws", webSocket)
	mux.HandleFunc("/", loadHome).Methods("GET")
	mux.HandleFunc("/login", login).Methods("POST")
	mux.HandleFunc("/chat", loadChat).Methods("GET")
	mux.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("front/css/"))))
	mux.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("front/js/"))))
	mux.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("front/img/"))))
	log.Fatalln(http.ListenAndServe(":80", mux))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func webSocket(res http.ResponseWriter, req *http.Request) {
	conn, error := upgrader.Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	client := &Client{id: uuid.NewV4().String(), socket: conn, send: make(chan []byte)}
	manager.register <- client
	go client.read()
	go client.write()
}

func loadHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=5")
	http.ServeFile(w, r, "front/index.html")
	log.Println("home")
}
func loadChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=5")
	log.Println("chat")
	http.ServeFile(w, r, "front/chat.html")
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("redirect")
	http.Redirect(w, r, "localhost/chat", 200)
}
