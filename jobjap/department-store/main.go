package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DepartmentStore struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Location      string `json:"location"`
	AvailableTime string `json:"available_time"`
	Distance      string `json:"distance"`
}

var department_stores = []DepartmentStore{
	{
		Name:          "CentralWorld",
		Description:   "One of the largest shopping malls in Thailand",
		Location:      "999/9 Rama I Rd, Pathum Wan, Bangkok",
		AvailableTime: "10:00 AM - 10:00 PM",
		Distance:      "0.1 KM",
	},
	{
		Name:          "Siam Paragon",
		Description:   "Luxury shopping mall with a wide range of international brands",
		Location:      "991 Rama I Rd, Pathum Wan, Bangkok",
		AvailableTime: "10:00 AM - 10:00 PM",
		Distance:      "1 KM",
	},
	{
		Name:          "MBK Center",
		Description:   "Popular mall known for affordable shopping options",
		Location:      "444 Phayathai Rd, Pathum Wan, Bangkok",
		AvailableTime: "10:00 AM - 10:00 PM",
		Distance:      "2 KM",
	},
	{
		Name:          "EmQuartier",
		Description:   "Trendy shopping center with international and local brands",
		Location:      "693, 695 Sukhumvit Rd, Khlong Tan Nuea, Bangkok",
		AvailableTime: "10:00 AM - 10:00 PM",
		Distance:      "1.3 KM",
	},
	{
		Name:          "Terminal 21",
		Description:   "Shopping mall with a unique airport theme and international zones",
		Location:      "88 Sukhumvit Rd, Khlong Toei, Bangkok",
		AvailableTime: "10:00 AM - 10:00 PM",
		Distance:      "1.4 KM",
	},
	{
		Name:          "ICONSIAM",
		Description:   "New riverside shopping destination with a mix of Thai and global brands",
		Location:      "299 Charoen Nakhon Rd, Khlong Ton Sai, Bangkok",
		AvailableTime: "10:00 AM - 10:00 PM",
		Distance:      "1.7 KM",
	},
	{
		Name:          "Central Embassy",
		Description:   "Upscale shopping mall with designer brands and high-end dining options",
		Location:      "1031 Ploenchit Rd, Pathum Wan, Bangkok",
		AvailableTime: "10:00 AM - 10:00 PM",
		Distance:      "1.9 KM",
	},
	{
		Name:          "The Mall Bangkapi",
		Description:   "Large shopping center with a variety of stores and entertainment options",
		Location:      "3522 Lat Phrao Rd, Khlong Chan, Bang Kapi, Bangkok",
		AvailableTime: "10:00 AM - 10:00 PM",
		Distance:      "1.1 KM",
	},
	{
		Name:          "Platinum Fashion Mall",
		Description:   "Famous for wholesale fashion and retail shopping",
		Location:      "222 Phetchaburi Rd, Thanon Phetchaburi, Ratchathewi, Bangkok",
		AvailableTime: "10:00 AM - 8:00 PM",
		Distance:      "3 KM",
	},
}

func filterStores(companyName string) []DepartmentStore {
	var startIndex int
	switch companyName {
	case "Agoda":
		startIndex = 0
	case "LineMan":
		startIndex = 3
	case "Shoppee":
		startIndex = 6
	default:
		return []DepartmentStore{}
	}

	endIndex := startIndex + 3
	if endIndex > len(department_stores) {
		endIndex = len(department_stores)
	}

	return department_stores[startIndex:endIndex]
}

func main() {

	router := gin.Default()
	router.GET("/department-store", func(c *gin.Context) {
		companyName := c.Query("companyName")
		filteredStores := filterStores(companyName)
		c.JSON(http.StatusOK, filteredStores)
	})

	router.Run(":8001")
}
