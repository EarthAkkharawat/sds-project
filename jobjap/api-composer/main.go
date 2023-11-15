package main

import (
	"encoding/json"
	"fmt"
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
type JobPosition struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Skill       string `json:"skill"`
	Salary      string `json:"salary"`
	Team        string `json:"team"`
}
type Parking struct {
	Name          string `json:"name"`
	Price         string `json:"price"`
	AvailableTime string `json:"available_time"`
	Location      string `json:"location"`
	Distance      string `json:"distance"`
	Slot          string `json:"slot"`
}

func getDepartmentStoreData(companyName string) ([]DepartmentStore, error) {
	// URL and client setup remains the same...
	path := "http://localhost:8001" + "/department-store?companyName=" + companyName
	println(string(path))

	// Create an HTTP client
	client := &http.Client{}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var jsonResponse []DepartmentStore
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}
func getJobPositionData(companyName string) ([]JobPosition, error) {
	// URL and client setup remains the same...
	path := "http://localhost:8002" + "/job-position?companyName=" + companyName

	// Create an HTTP client
	client := &http.Client{}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var jsonResponse []JobPosition
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}

func getParkingData(companyName string) ([]Parking, error) {
	path := "http://localhost:8003" + "/parking?companyName=" + companyName
	// Create an HTTP client
	client := &http.Client{}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var jsonResponse []Parking
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		companyName := c.Query("companyName")
		departmentStoreData, err1 := getDepartmentStoreData(companyName)
		if err1 != nil {
			fmt.Println(err1)
		}
		parkingData, err2 := getParkingData(companyName)
		if err2 != nil {
			fmt.Println(err2)
		}
		jobPositionData, err3 := getJobPositionData(companyName)
		if err3 != nil {
			fmt.Println(err3)
		}

		c.JSON(http.StatusOK, gin.H{
			"companyName":         companyName,
			"departmentStoreData": departmentStoreData,
			"parkingData":         parkingData,
			"jobPositionData":     jobPositionData,
		})
	})
	router.Run(":8000")
}
