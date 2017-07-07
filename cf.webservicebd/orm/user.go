package orm

import "time"

type User struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Createdtime time.Time `json:"created_time" gorm:"column:created_time"`
	Editedtime  time.Time `json:"edited_time" gorm:"type:timestamp;column:edited_time"`
}

type Users []User

func (this User) Create(username, password, email string) {

	this.Username = username
	this.Password = password
	this.Email = email
	this.Createdtime = time.Now()
	db.Create(&this)
}

func (this *User) Find(id *int64) *User {
	db.Where("id=?", id).First(this)
	return this
}

func (this *Users) FindAll() *Users {
	db.Find(this)
	return this
}

func (this *User) Update() {
	db.Model(&this).UpdateColumns(&this)
}

func (this *User) Delete() {
	db.Delete(this)
}
//gg