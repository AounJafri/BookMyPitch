package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Aounjafri/BookMyPitch/Backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func BookGround(c *gin.Context) {
	var groundid int

	if currentLoggedInUserId != 0 && currentUserName != "" {

		var bookingInfo models.Booking
		id := c.Param("id")
		groundid, _ = strconv.Atoi(id)

		bookingInfo.UserId = currentLoggedInUserId
		bookingInfo.GroundId = groundid
		bookingInfo.Status = "Awaiting approval"

		date := c.Request.FormValue("date")
		timeslot := c.Request.FormValue("timeSlot")

		bookingInfo.BookingDate = date
		bookingInfo.Timeslot = timeslot

		fmt.Println(bookingInfo)

		isBooked := models.MakeBooking(bookingInfo)

		if isBooked {
			c.HTML(400, "book.html", gin.H{
				"Error":    "Slot already Booked",
				"GroundId": groundid,
				"Today":    time.Now().Format("2006-01-02"),
			})
		} else {
			c.HTML(200, "book.html", gin.H{
				"Success":  "Booking request submitted successfully",
				"GroundId": groundid,
				"Today":    time.Now().Format("2006-01-02"),
			})
			return
		}

	} else {
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func GetBookings(c *gin.Context) {
	url := c.Request.URL.String()

	if currentLoggedInUserId != 0 && currentUserName != "" {

		user := models.GetUserByName(currentUserName)

		if url == "/bookings" {
			if user.Role == "Admin" {

				bookings := models.GetBookings()
				fmt.Println(bookings)
				c.HTML(200, "bookings.html", gin.H{
					"userRole": "Admin",
					"Bookings": bookings,
				})
				return
			} else {
				bookings := models.GetUserBookings(currentLoggedInUserId)
				fmt.Println(bookings)
				c.HTML(200, "bookings.html", gin.H{
					"userRole": "User",
					"Bookings": bookings,
				})
				return
			}
		}
		id := c.Param("id")
		groundid, _ := strconv.Atoi(id)

		c.HTML(200, "book.html", gin.H{
			"GroundId": groundid,
			"Today":    time.Now().Format("2006-01-02"),
		})
		return

	} else {
		c.Redirect(http.StatusSeeOther, "/")

	}
}

func ApproveBooking(c *gin.Context) {
	id := c.Param("bookingid")
	bookingid, _ := strconv.Atoi(id)

	if currentLoggedInUserId != 0 && currentUserName != "" {

		models.ApproveBooking(bookingid)

		c.Redirect(http.StatusSeeOther, "/bookings")

	} else {
		c.Redirect(http.StatusSeeOther, "/")

	}
}
