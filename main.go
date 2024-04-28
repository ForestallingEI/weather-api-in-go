package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"weatherapi/structs"
)

// API のベースURL
const baseURL = "https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=yes"


// 天気情報を取得
func getWeather(city string) (structs.Weather, error) {
	url := fmt.Sprintf(baseURL, os.Getenv("WEATHER_API_KEY"), city)
	resp, err := http.Get(url)
	if err != nil {
		return structs.Weather{}, err
	}
	defer resp.Body.Close()

	var data structs.Weather
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return structs.Weather{}, err
	}
	return data, nil
}

// ルートハンドラ
func rootHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getWeather("London")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")
	data, err := getWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("weather_result.html"))
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/weather", weatherHandler)
	fmt.Println("サーバー起動: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
