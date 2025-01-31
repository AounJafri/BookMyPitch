package models

import (
	"database/sql"
	"fmt"

	"github.com/Aounjafri/BookMyPitch/Backend/pkg/config"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Role     string `json:"role"`
	Location string `json:"location"`
}

var db *sql.DB

func init() {
	db = config.Connect()
}

func CreateUser(user User) {
	_, err := db.Exec("INSERT INTO Users(username,password,email,age,role,location) Values(?,?,?,?,?,?)", user.Username, user.Password, user.Email, user.Age, user.Role, user.Location)

	if err != nil {
		fmt.Println(err.Error())
	}

}

func GetUserByName(username string) User {
	var user User

	db.QueryRow("Select id,password,role From Users Where username=?", username).Scan(&user.Id, &user.Password, &user.Role)

	return user
}

func GetUsers() []User {
	results, err := db.Query("Select * From Users Where role = ?", "User")
	if err != nil {
		fmt.Println(err.Error())
	}
	var users []User

	for results.Next() {
		var user User
		results.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Age, &user.Role, &user.Location)
		users = append(users, user)
	}
	return users
}

func DeleteUser(id int) {
	_, er := db.Exec("Delete from Bookings where userid = ?", id)
	if er != nil {
		fmt.Println(er.Error())
	}
	_, err := db.Exec("Delete From Users Where id = ?", id)

	if err != nil {
		fmt.Println(err.Error())
	}
}
