package models

import "errors"

//User es la estructura que contiene los campos del usuario
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const userSchema string = `CREATE TABLE users (
  id int(6) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  password varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  email varchar(40) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  create_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP) 
  ENGINE=Aria DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci PAGE_CHECKSUM=1;`

type Users []User

var users = make(map[int]User)

func SetDefaultUser() {
	user := User{ID: 1, Username: "Arimdor", Password: "password123"}
	users[user.ID] = user
}

func GetUsers() Users {
	list := Users{}
	for _, user := range users {
		list = append(list, user)
	}
	return list
}

func GetUser(userID int) (User, error) {
	if user, ok := users[userID]; ok {
		return user, nil
	}
	return User{}, errors.New("No se encontro el usuario solicitado")
}

func SaveUser(user User) User {
	user.ID = len(users) + 1
	users[user.ID] = user
	return user
}

func UpdateUser(user User, username string, password string) User {
	user.Username = username
	user.Password = password
	users[user.ID] = user
	return user
}

func DeleteUser(id int) {
	delete(users, id)
}
