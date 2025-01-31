package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Aounjafri/BookMyPitch/Backend/pkg/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var currentLoggedInUserId int
var currentUserName string

func HomePage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"Title": "BookMyPitch",
	})
}

// func Form(c *gin.Context) {
// 	url := c.Request.URL.String()
// 	fmt.Println(url)
// 	fmt.Println(c.Request.Method)
// 	fmt.Println(c.Request.Method == "GET")
// 	if url == "/admin" {
// 		c.HTML(200, "form.html", gin.H{})

// 	} else if url == "/login" {

// 	} else if url == "/register" {

// 	}
// }

func Register(c *gin.Context) {

	if c.Request.Method == "GET" {
		url := c.Request.URL.String()
		if url == "/register" {
			c.HTML(200, "form.html", gin.H{
				"Title":       "User Registeration",
				"Heading":     "Register Now",
				"Subheading":  "Sign up to create your account",
				"Type":        "registration",
				"FormAction":  "/register",
				"ButtonLabel": "Register",
			})
		}
		return
	}

	var user models.User

	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	user.Email = c.Request.FormValue("email")
	user.Age, _ = strconv.Atoi(c.Request.FormValue("age"))
	user.Location = c.Request.FormValue("location")

	fmt.Println(user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	user.Password = string(hashedPassword)
	user.Role = "User"

	models.CreateUser(user)

	c.HTML(200, "form.html", gin.H{
		"Title":       "User Registeration",
		"Heading":     "Register Now",
		"Subheading":  "Sign up to create your account",
		"Type":        "registration",
		"FormAction":  "/register",
		"ButtonLabel": "Register",
		"Error":       "Registeration Successfull, Login Now to access your account",
	})

}

func Login(c *gin.Context) {

	url := c.Request.URL.String()

	if c.Request.Method == "GET" {
		if url == "/admin" {
			c.HTML(200, "form.html", gin.H{
				"Title":       "Admin Login",
				"Heading":     "Welcome Admin",
				"Subheading":  "Please enter your credentials to login",
				"Type":        "admin-login",
				"FormAction":  "/admin/login",
				"ButtonLabel": "Login",
			})
		} else if url == "/login" {
			c.HTML(200, "form.html", gin.H{
				"Title":       "User Login",
				"Heading":     "User Login",
				"Subheading":  "Please enter your credentials to login",
				"Type":        "login",
				"FormAction":  "/login",
				"ButtonLabel": "Login",
			})
		}
		return
	}

	//SIMPLE USER LOGIN FUNCTIONALITY

	if url == "/login" {

		username := c.Request.FormValue("username")
		password := c.Request.FormValue("password")

		storedUser := models.GetUserByName(username)
		fmt.Println(storedUser)

		if storedUser.Id == 0 {
			c.HTML(400, "form.html", gin.H{
				"Title":       "User Login",
				"Heading":     "User Login",
				"Subheading":  "Please enter your credentials to login",
				"Type":        "login",
				"FormAction":  "/login",
				"ButtonLabel": "Login",
				"Error":       "User does not exists",
			})
			return
		}
		fmt.Println(password)
		err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password))
		fmt.Println(err)
		if err != nil {
			c.HTML(400, "form.html", gin.H{
				"Title":       "User Login",
				"Heading":     "User Login",
				"Subheading":  "Please enter your credentials to login",
				"Type":        "login",
				"FormAction":  "/login",
				"ButtonLabel": "Login",
				"Error":       "Password Incorrect",
			})
			return
		}

		currentLoggedInUserId = storedUser.Id
		currentUserName = username

		fmt.Println(currentLoggedInUserId)

		// c.HTML(200, "grounds.html", gin.H{
		// 	"Title":       "User Login",
		// 	"Heading":     "User Login",
		// 	"Subheading":  "Please enter your credentials to login",
		// 	"Type":        "login",
		// 	"FormAction":  "/login",
		// 	"ButtonLabel": "Login",
		// 	"Error":       "Login Successfull, Will redirect you to home page in a bit",
		// })
		c.Redirect(http.StatusMovedPermanently, "/grounds")

		//ADMIN LOGIN FUNCTIONALITY

	} else if url == "/admin/login" {

		username := c.Request.FormValue("username")
		password := c.Request.FormValue("password")

		storedUser := models.GetUserByName(username)
		fmt.Println(storedUser)

		if storedUser.Id == 0 {
			c.HTML(400, "form.html", gin.H{
				"Title":       "Admin Login",
				"Heading":     "Welcome Admin",
				"Subheading":  "Please enter your credentials to login",
				"Type":        "admin-login",
				"FormAction":  "/admin/login",
				"ButtonLabel": "Login",
				"Error":       "Admin account does not exists",
			})
			return
		}
		if storedUser.Role != "Admin" {
			c.HTML(400, "form.html", gin.H{
				"Title":       "Admin Login",
				"Heading":     "Welcome Admin",
				"Subheading":  "Please enter your credentials to login",
				"Type":        "admin-login",
				"FormAction":  "/admin/login",
				"ButtonLabel": "Login",
				"Error":       "This is not admin account",
			})
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password))

		if err != nil {
			c.HTML(400, "form.html", gin.H{
				"Title":       "User Login",
				"Heading":     "User Login",
				"Subheading":  "Please enter your credentials to login",
				"Type":        "login",
				"FormAction":  "/login",
				"ButtonLabel": "Login",
				"Error":       "Invalid admin credentials",
			})
			return
		}

		currentLoggedInUserId = storedUser.Id
		currentUserName = username

		fmt.Println(currentLoggedInUserId)

		c.Redirect(http.StatusMovedPermanently, "/grounds")
	}
}

func Logout(c *gin.Context) {

	if currentLoggedInUserId != 0 && currentUserName != "" {

		currentLoggedInUserId = 0
		currentUserName = ""
		c.JSON(200, "Logout Successfull")
	} else {
		c.JSON(400, "No user is logged in")
	}
}

func GetUsers(c *gin.Context) {
	if currentLoggedInUserId != 0 && currentUserName != "" {
		users := models.GetUsers()
		c.HTML(200, "users.html", gin.H{
			"Users": users,
		})

	} else {
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func DeleteUser(c *gin.Context) {
	if currentLoggedInUserId != 0 && currentUserName != "" {
		id := c.Param("id")
		idi, _ := strconv.Atoi(id)

		models.DeleteUser(idi)

		c.Redirect(http.StatusSeeOther, "/users")

	} else {
		c.Redirect(http.StatusSeeOther, "/")
	}
}
