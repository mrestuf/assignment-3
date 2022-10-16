package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		for range time.Tick(time.Second * 15) {
			WriteJSONFile()
		}
	}()
	startRoute := StartApp()
	startRoute.Run(":4000")
}

func StartApp() *gin.Engine {
	r := gin.Default()
	r.GET("/index", GetStatusHandler)
	return r
}

type StatusJSON struct {
	Status map[string]int `json: "status"`
}

type Status struct {
	Water       int    `json:"water"`
	WaterStatus string `json:"waterStatus"`
	Wind        int    `json:"wind"`
	WindStatus  string `json:"windStatus"`
}

func WriteJSONFile() {
	data := StatusJSON{
		Status: map[string]int{
			"water": rand.Intn(100),
			"wind":  rand.Intn(100),
		},
	}

	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("status.json", jsonData, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func GetStatus() Status {
	var status Status
	statusJSON := StatusJSON{}
	data, err := ioutil.ReadFile("status.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &statusJSON)
	if err != nil {
		fmt.Println(err)
	}

	water := statusJSON.Status["water"]
	wind := statusJSON.Status["wind"]

	if water <= 5 {
		status.WaterStatus = "Aman"
	} else if 6 <= water && water <= 8 {
		status.WaterStatus = "Siaga"
	} else {
		status.WaterStatus = "Bahaya"
	}

	if wind <= 6 {
		status.WindStatus = "Aman"
	} else if 7 <= wind && wind <= 15 {
		status.WindStatus = "Siaga"
	} else {
		status.WindStatus = "Bahaya"
	}

	status.Water = water
	status.Wind = wind

	// ctx.JSON(http.StatusOK, status)

	return status
}

func GetStatusHandler(ctx *gin.Context) {
	data := GetStatus()

	var t, err = template.ParseFiles("page/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(ctx.Writer, data)
}
