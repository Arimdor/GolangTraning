package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../models"
	"../orm"
	"github.com/gorilla/mux"
)

// GetUsers devuelve todos los usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users orm.Users
	models.SendData(w, users.FindAll())
}

// GetUser devuelve un usuario mediante el ID
func GetUser(w http.ResponseWriter, r *http.Request) {

	if user, err := getUserByRequest(r); err != nil {
		w.WriteHeader(http.StatusNotFound)
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}
}

// CreateUser crea un usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if r.Body == nil {
		http.Error(w, "Envio vacio", 400)
		log.Fatal("Envio vacio")
		return
	}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
		log.Println("Error: ", err.Error())
		log.Println(user)
	} else {
		models.SendData(w, models.SaveUser(user))
	}
}

// UpdateUser actualiza un usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	userResponse := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	user = models.UpdateUser(user, userResponse.Username, userResponse.Password)
	models.SendData(w, user)

}

// DeleteUser elimina un usuario
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	models.DeleteUser(user.ID)
	models.SendNotContent(w)
}

func getUserByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	user, err := models.GetUser(userID)
	if err != nil {
		return user, err
	}
	return user, nil
}
//gg