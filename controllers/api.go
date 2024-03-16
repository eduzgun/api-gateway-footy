package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/eduzgun/api-gateway-footy/utils"
	"github.com/gin-gonic/gin"
)

type Statistic struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type FixtureStatisticResponse struct {
	Team       Team        `json:"team"`
	Statistics []Statistic `json:"statistics"`
}

type APIResponse struct {
	Get      string                     `json:"get"`
	Errors   []interface{}              `json:"errors"`
	Results  int                        `json:"results"`
	Paging   map[string]int             `json:"paging"`
	Response []FixtureStatisticResponse `json:"response"`
}

func Fixture(c *gin.Context) {
	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorised"})
		return
	}

	if claims.Role != "user" && claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorised"})
		return
	}

	//now make a request to rapidapi
	url := "https://api-football-v1.p.rapidapi.com/v3/fixtures/statistics?fixture=1035442"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "76a5c68254msh15c95aa5f37d156p1a408ajsn7bc8944039ed")
	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	//fmt.Println(res)
	//fmt.Println(body)
	var apiResponse APIResponse
	err = json.Unmarshal([]byte(body), &apiResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data []map[string]interface{}
	for _, fixture := range apiResponse.Response {
		fixtureData := make(map[string]interface{})
		fixtureData["team"] = fixture.Team
		fixtureData["statistics"] = fixture.Statistics
		data = append(data, fixtureData)
		fmt.Println(data)
	}
	c.JSON(200, gin.H{"success": "home page", "role": claims.Role, "content": data})
}
