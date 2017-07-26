package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Response struct {
	Mensaje string `json:"messaje"`
	Status  int    `json:"status"`
	IsValid bool   `json:"isvalid"`
}

var Users = struct {
	sync.RWMutex
	m map[string]User
}{m: make(map[string]User)}

type User struct {
	Username  string
	WebSocket *websocket.Conn
}

func createUser(username string, ws *websocket.Conn) User {
	return User{username, ws}
}

func addUser(user User) {
	Users.Lock()
	defer Users.Unlock()
	log.Println(user)
	Users.m[user.Username] = user
}

func removeUser(username string) {
	Users.Lock()
	defer Users.Unlock()
	delete(Users.m, username)
	log.Println("El usuario se fue")
}

func userExist(username string) bool {
	Users.RLock()
	defer Users.RUnlock()
	if _, ok := Users.m[username]; ok {
		log.Println("existe")
		return true
	}
	log.Println("no existe")
	return false
}

func toArryByte(value string) []byte {
	return []byte(value)
}

func concatMessage(username string, arreglo []byte) string {
	return username + " : " + string(arreglo[:])
}

func sendMessage(type_message int, message []byte) {
	Users.RLock()
	defer Users.RUnlock()
	for _, user := range Users.m {
		if err := user.WebSocket.WriteMessage(type_message, message); err != nil {
			log.Println(err)
			return
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	//rutas web
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/validate", validate).Methods("POST")
	router.HandleFunc("/chat/{username}", webSocket).Methods("GET")
	router.HandleFunc("/index", loadStatic).Methods("GET")
	router.HandleFunc("/", loadStatic).Methods("GET")

	//archivos staticos
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("front/css/"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("front/js/"))))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("front/img/"))))

	//levantando servidor
	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:80",
		// WriteTimeout: 500 * time.Second,
		// ReadTimeout:  500 * time.Second,
	}
	log.Println("El servidor esta a la escucha 80")
	log.Fatal(srv.ListenAndServe())
}

func loadStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "front/index.html")
}

func validate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(&Users)
	username := r.FormValue("username")
	response := Response{}
	log.Println(username)
	if userExist(username) {
		log.Println("el usuario ya existe")
		response.IsValid = false
	} else {
		log.Println("Usuario Valido")
		response.IsValid = true
	}
	json.NewEncoder(w).Encode(response)

}

func webSocket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	currentuser := createUser(username, ws)
	addUser(currentuser)
	log.Println("Usuario agregado al mapa")
	for {
		typemessage, message, err := ws.ReadMessage()
		if err != nil {
			removeUser(username)
			return
		}
		finalmessage := concatMessage(username, message)
		sendMessage(typemessage, toArryByte(finalmessage))
	}
}

func createResponse(messaje string, status int, valid bool) Response {
	return Response{messaje, status, valid}
}

// func holaMundo(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hola Mundo"))
// }
// func holajson(w http.ResponseWriter, r *http.Request) {
// 	response := createResponse()
// 	json.NewEncoder(w).Encode(response)
// }
