package models

import "fmt"

type Ground struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	PricePerDay int    `json:"priceperday"`
	Image       string `json:"image"`
}

func AddGround(ground Ground) {
	_, err := db.Exec("Insert into Grounds(name,location,priceperday,image) Values(?,?,?,?)", ground.Name, ground.Location, ground.PricePerDay, ground.Image)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func GetGroundById(id int) Ground {
	var ground Ground

	db.QueryRow("Select * From Grounds Where id=?", id).Scan(&ground.Id, &ground.Name, &ground.Location, &ground.PricePerDay, &ground.Image)
	return ground
}

func GetGrounds() []Ground {
	var grounds []Ground

	results, err := db.Query("Select * From Grounds")
	if err != nil {
		fmt.Println(err.Error())
	}

	for results.Next() {
		var gr Ground

		results.Scan(&gr.Id, &gr.Name, &gr.Location, &gr.PricePerDay, &gr.Image)
		grounds = append(grounds, gr)
	}
	return grounds
}

func UpdateGround(id int, updatedGr Ground) Ground {
	storedGround := GetGroundById(id)

	if updatedGr.Name != storedGround.Name && updatedGr.Name != "" {
		storedGround.Name = updatedGr.Name
	}
	if updatedGr.Location != storedGround.Location && updatedGr.Location != "" {
		storedGround.Location = updatedGr.Location
	}
	if updatedGr.PricePerDay != storedGround.PricePerDay && updatedGr.PricePerDay != 0 {
		storedGround.PricePerDay = updatedGr.PricePerDay
	}

	_, err := db.Exec("Update Grounds Set name=?, location=?, priceperday=? Where id = ?", storedGround.Name, storedGround.Location, storedGround.PricePerDay, id)

	if err != nil {
		fmt.Println(err.Error())
	}
	return storedGround
}

func DeleteGround(id int) error {
	_, err := db.Exec("Delete From Grounds Where id=?", id)

	return err
}
