package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Aounjafri/BookMyPitch/Backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func AddGround(c *gin.Context) {

	if currentLoggedInUserId != 0 && currentUserName != "" {

		if c.Request.Method == "GET" {
			c.HTML(200, "add-ground.html", nil)
			return
		}

		var ground models.Ground

		ground.Name = c.Request.FormValue("name")
		ground.Location = c.Request.FormValue("location")
		ground.PricePerDay, _ = strconv.Atoi(c.Request.FormValue("price"))
		ground.Image = c.Request.FormValue("image")

		fmt.Println(ground)

		models.AddGround(ground)

		c.Redirect(http.StatusSeeOther, "/grounds")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}

}

func GetGrounds(c *gin.Context) {

	if currentLoggedInUserId != 0 && currentUserName != "" {

		user := models.GetUserByName(currentUserName)

		grounds := models.GetGrounds()

		if len(grounds) == 0 {
			if user.Role == "Admin" {
				c.HTML(200, "grounds.html", gin.H{
					"userRole": "Admin",
				})
				return
			} else {
				c.HTML(200, "grounds.html", gin.H{
					"userRole": "User",
				})
			}
		} else {
			if user.Role == "Admin" {
				c.HTML(200, "grounds.html", gin.H{
					"userRole": "Admin",
					"grounds":  grounds,
				})
				return
			} else {
				c.HTML(200, "grounds.html", gin.H{
					"userRole": "User",
					"grounds":  grounds,
				})
			}
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func DeleteGround(c *gin.Context) {
	id := c.Param("id")

	inid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err.Error())
	}

	if currentLoggedInUserId != 0 && currentUserName != "" {

		err := models.DeleteGround(inid)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		c.Redirect(http.StatusSeeOther, "/grounds")

	} else {
		c.Redirect(http.StatusSeeOther, "/")
	}
}
