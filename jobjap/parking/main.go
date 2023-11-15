package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Parking struct {
	Name          string `json:"name"`
	Price         string `json:"price"`
	AvailableTime string `json:"available_time"`
	Location      string `json:"location"`
	Distance      string `json:"distance"`
	Slot          string `json:"slot"`
}

var parkings = []Parking{
	{
		Name:          "BenjaKit Parking",
		Price:         "100 THB/h",
		AvailableTime: "24 hours",
		Location:      "Pathumwan",
		Distance:      "0.5 KM",
		Slot:          "57",
	},
	{
		Name:          "Payathai Plaza Parking",
		Price:         "80 THB/h",
		AvailableTime: "7am - 11pm",
		Location:      "Payathai",
		Distance:      "0.3 KM",
		Slot:          "33",
	},
	{
		Name:          "Sukhumvit Square Parking",
		Price:         "120 THB/h",
		AvailableTime: "24 hours",
		Location:      "Sukhumvit",
		Distance:      "0.5 KM",
		Slot:          "432",
	},
	{
		Name:          "100year Garden Parking",
		Price:         "90 THB/h",
		AvailableTime: "9am - 10pm",
		Location:      "Pathumwan",
		Distance:      "0.8 KM",
		Slot:          "234",
	},
	{
		Name:          "BTS Payathai Station Parking",
		Price:         "70 THB/h",
		AvailableTime: "6am - 12am",
		Location:      "Payathai",
		Distance:      "0.1 KM",
		Slot:          "103",
	},
	{
		Name:          "Airport Link Parking",
		Price:         "110 THB/h",
		AvailableTime: "8am - 11pm",
		Location:      "Sukhumvit",
		Distance:      "0.9 KM",
		Slot:          "100",
	},
}

func filterParkings(companyName string) []Parking {
	var startIndex int
	switch companyName {
	case "Agoda":
		startIndex = 0
	case "LineMan":
		startIndex = 2
	case "Shoppee":
		startIndex = 4
	default:
		return []Parking{}
	}

	endIndex := startIndex + 2
	if endIndex > len(parkings) {
		endIndex = len(parkings)
	}

	return parkings[startIndex:endIndex]
}

func main() {

	router := gin.Default()
	router.GET("/parking", func(c *gin.Context) {
		companyName := c.Query("companyName")
		filteredParkings := filterParkings(companyName)
		c.JSON(http.StatusOK, filteredParkings)
	})

	router.Run(":8003")
}
