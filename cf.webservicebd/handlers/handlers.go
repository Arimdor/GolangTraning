package handlers

import (
	"net/http"
	"strconv"

	"../models"
	"github.com/julienschmidt/httprouter"
)

// GetAllProducts devuelve todos los usuarios
func GetAllProducts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var products models.Products
	products.FindAll()
	models.SendData(w, products)
}

// GetProduct devuelve un usuario mediante el ID
func GetProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var product models.Product
	id, _ := strconv.Atoi(p.ByName("id"))
	product.Find(id)
	models.SendData(w, product)
}

// // CreateUser crea un usuario
// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	user := models.User{}
// 	if r.Body == nil {
// 		http.Error(w, "Envio vacio", 400)
// 		log.Fatal("Envio vacio")
// 		return
// 	}
// 	decoder := json.NewDecoder(r.Body)

// 	if err := decoder.Decode(&user); err != nil {
// 		models.SendUnprocessableEntity(w)
// 		log.Println("Error: ", err.Error())
// 		log.Println(user)
// 	} else {
// 		models.SendData(w, models.SaveUser(user))
// 	}
// }

// // UpdateUser actualiza un usuario
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	user, err := getUserByRequest(r)
// 	if err != nil {
// 		models.SendNotFound(w)
// 		return
// 	}
// 	userResponse := models.User{}
// 	decoder := json.NewDecoder(r.Body)

// 	if err := decoder.Decode(&userResponse); err != nil {
// 		models.SendUnprocessableEntity(w)
// 		return
// 	}
// 	user = models.UpdateUser(user, userResponse.Username, userResponse.Password)
// 	models.SendData(w, user)

// }

// // DeleteUser elimina un usuario
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	user, err := getUserByRequest(r)
// 	if err != nil {
// 		models.SendNotFound(w)
// 		return
// 	}
// 	models.DeleteUser(user.ID)
// 	models.SendNotContent(w)
// }
