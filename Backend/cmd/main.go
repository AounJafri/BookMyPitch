package main

import (
	"fmt"

	"github.com/Aounjafri/BookMyPitch/Backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r gin.IRoutes) {

	//Registeration and Login Routes
	r.GET("/", controllers.HomePage)
	r.GET("/login", controllers.Login)
	r.GET("/admin", controllers.Login)
	r.GET("/register", controllers.Register)

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/admin/login", controllers.Login)
	r.POST("logout", controllers.Logout)

	// // Admin Handlers
	r.GET("/users", controllers.GetUsers)
	r.POST("/user/delete/:id", controllers.DeleteUser)
	r.POST("/ground", controllers.AddGround)
	r.GET("/ground/delete/:id", controllers.DeleteGround)
	r.GET("/ground", controllers.AddGround)
	r.POST("/admin/approve/booking/:bookingid", controllers.ApproveBooking)

	// User Handlers
	r.GET("/grounds", controllers.GetGrounds)
	r.POST("/ground/book/:id", controllers.BookGround)
	r.GET("/ground/book/:id", controllers.GetBookings)
	r.GET("/bookings", controllers.GetBookings)
}

func main() {
	r := gin.Default()
	staticPath := "../../Frontend/static"
	templatesPath := "../../Frontend/templates/*"

	r.Static("/static", staticPath)

	r.LoadHTMLGlob(templatesPath)
	Routes(r)

	fmt.Println("Starting server at port 8080")
	r.Run(":8080")

}
