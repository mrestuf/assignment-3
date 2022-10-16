package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/index", GetStatusHandler)

	log.Println("server running at port", ":4000")
	http.ListenAndServe(":4000", nil)
}

type Status struct {
	Water       int
	WaterStatus string
	Wind        int
	WindStatus  string
}

func GetStatus() Status {
	var status Status
	rand.Seed(time.Now().Unix())

	status.Water = rand.Intn(100) + 1
	status.Wind = rand.Intn(100) + 1

	if status.Water <= 5 {
		status.WaterStatus = "Aman"
	} else if status.Water > 5 && status.Water < 9 {
		status.WaterStatus = "Siaga"
	} else if status.Water >= 9 {
		status.WaterStatus = "Bahaya"
	} else {
		status.WaterStatus = "-"
	}

	if status.Wind <= 6 {
		status.WindStatus = "Aman"
	} else if status.Wind > 6 && status.Wind < 16 {
		status.WindStatus = "Siaga"
	} else if status.Wind >= 16 {
		status.WindStatus = "Bahaya"
	} else {
		status.WindStatus = "-"
	}

	return status
}

func GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	data := GetStatus()

	tmpl, err := template.ParseFiles("./page/index.html")
	if err != nil {
		fmt.Fprintf(w, "error")
		return
	}

	tmpl.Execute(w, data)
}
