package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobPosition struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Skill       string `json:"skill"`
	Salary      string `json:"salary"`
	Team        string `json:"team"`
}

var positions = []JobPosition{
	{
		Name:        "Software Engineer",
		Description: "Responsible for developing and maintaining software applications.",
		Skill:       "Proficient in programming languages like Go, Python, or Java.",
		Salary:      "$80,000 - $120,000",
		Team:        "Engineering",
	},
	{
		Name:        "Product Manager",
		Description: "Oversees the development and lifecycle of a product.",
		Skill:       "Strong project management and communication skills.",
		Salary:      "$90,000 - $130,000",
		Team:        "Product",
	},
	{
		Name:        "Data Scientist",
		Description: "Analyzes complex data sets to derive actionable insights.",
		Skill:       "Expertise in data analysis and machine learning.",
		Salary:      "$85,000 - $125,000",
		Team:        "Data Analysis",
	},
	{
		Name:        "UX/UI Designer",
		Description: "Designs and improves user interfaces and experiences.",
		Skill:       "Proficient in design tools and user experience best practices.",
		Salary:      "$70,000 - $110,000",
		Team:        "Design",
	},
	{
		Name:        "DevOps Engineer",
		Description: "Manages infrastructure, deployment pipelines, and cloud services.",
		Skill:       "Experienced in cloud platforms, CI/CD, and infrastructure automation.",
		Salary:      "$90,000 - $130,000",
		Team:        "Operations",
	},
	{
		Name:        "Quality Assurance Engineer",
		Description: "Ensures the quality of software through manual and automated testing.",
		Skill:       "Experience in testing methodologies and tools.",
		Salary:      "$65,000 - $95,000",
		Team:        "Quality Assurance",
	},
	{
		Name:        "Technical Support Specialist",
		Description: "Provides technical support and troubleshoots issues for users.",
		Skill:       "Strong problem-solving skills and technical knowledge.",
		Salary:      "$50,000 - $70,000",
		Team:        "Support",
	},
	{
		Name:        "Cybersecurity Analyst",
		Description: "Protects company's data and infrastructure from cyber threats.",
		Skill:       "Knowledgeable in security protocols and threat detection.",
		Salary:      "$75,000 - $105,000",
		Team:        "Security",
	},
	{
		Name:        "Marketing Specialist",
		Description: "Develops and implements marketing strategies.",
		Skill:       "Expertise in digital marketing and analytics.",
		Salary:      "$60,000 - $85,000",
		Team:        "Marketing",
	},
}

func filterPositions(companyName string) []JobPosition {
	var startIndex int
	switch companyName {
	case "Agoda":
		startIndex = 0
	case "LineMan":
		startIndex = 3
	case "Shoppee":
		startIndex = 6
	default:
		return []JobPosition{}
	}

	endIndex := startIndex + 3
	if endIndex > len(positions) {
		endIndex = len(positions)
	}

	return positions[startIndex:endIndex]
}

func main() {

	router := gin.Default()
	router.GET("/job-position", func(c *gin.Context) {
		companyName := c.Query("companyName")
		filteredPositions := filterPositions(companyName)
		c.JSON(http.StatusOK, filteredPositions)
	})

	router.Run(":8002")
}
