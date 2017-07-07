package models

//User es la estructura que contiene los campos del usuario
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

const userSchema string = `CREATE TABLE users (
  id int(6) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  password varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  email varchar(40) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  create_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP) 
  ENGINE=Aria DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci PAGE_CHECKSUM=1;`

type Users []User

func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

func (this *User) Save() {
	if this.ID == 0 {
		this.instert()
	} else {
		this.update()
	}
}

func (this *User) instert() {
	sql := "INSERT users set username=?, password=?, email=?"
	result, _ := Exec(sql, this.Username, this.Password, this.Email)
	this.ID, _ = result.LastInsertId()
}

func (this *User) update() {
	sql := "UPDATE users SET username=?, password=?, email=?"
	Exec(sql, this.Username, this.Password, this.Email)
}

func (this *User) Delete() {
	sql := "DELETE FROM users where id=?"
	Exec(sql, this.ID)
}

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

func GetUser(id int) *User {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := Query(sql, id)
	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	}

	return user
}

func GetUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, _ := Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}
