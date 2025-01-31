package models

import (
	"fmt"
)

type Booking struct {
	Id          int    `json:"id"`
	GroundId    int    `json:"groundid"`
	UserId      int    `json:"userid"`
	BookingDate string `json:"bookingdate"`
	Timeslot    string `json:"timeslot"`
	Status      string `json:"status"`
}

func MakeBooking(booking Booking) bool {
	var count int
	err := db.QueryRow("Select Count(*) From Bookings where groundid = ? AND bookingdate = ? AND timeslot = ? ", booking.GroundId, booking.BookingDate, booking.Timeslot).Scan(&count)

	if err != nil {
		fmt.Println("Checking availability error :  " + err.Error())
	}

	if count > 0 {
		return true
	}

	_, errr := db.Exec("Insert into Bookings (groundid,userid,bookingdate,timeslot,status) Values(?,?,?,?,?)", booking.GroundId, booking.UserId, booking.BookingDate, booking.Timeslot, booking.Status)
	if errr != nil {
		fmt.Println("Inserting error:  " + errr.Error())
		return true
	}

	return false
}

func GetBookings() []Booking {
	var bookings []Booking

	results, err := db.Query("Select * From Bookings")
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {
		var bk Booking

		results.Scan(&bk.Id, &bk.GroundId, &bk.UserId, &bk.BookingDate, &bk.Timeslot, &bk.Status)
		bookings = append(bookings, bk)
	}
	return bookings
}

func GetUserBookings(id int) []Booking {

	var bookings []Booking

	results, err := db.Query("Select * From Bookings Where userid=?", id)
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {
		var bk Booking

		results.Scan(&bk.Id, &bk.GroundId, &bk.UserId, &bk.BookingDate, &bk.Timeslot, &bk.Status)
		bookings = append(bookings, bk)
	}
	return bookings
}

func ApproveBooking(bookingid int) {
	_, err := db.Exec("Update Bookings set status = ? where id = ?", "Approved", bookingid)

	if err != nil {
		fmt.Println("Updation error :  " + err.Error())
	}

}
