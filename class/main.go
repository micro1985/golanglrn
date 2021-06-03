package main

import ( 
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"flag"
)

const (
	basic_url = "http://api.weatherapi.com/v1/current.json?key=34976b7e974d4c67beb163221210106&q="
)

type RequestData struct{
	city string
	air string
}

type GotData struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

func build_url(r RequestData) string {
	res := basic_url + r.city + "&aqi=" + r.air
	return res
}

func parse_request(str string) {
	data := GotData{}
    json.Unmarshal([]byte(str), &data)
    fmt.Printf("Country: %s", data.Location.Country)
	fmt.Println("\n")
	fmt.Printf("Name: %s", data.Location.Name)
	fmt.Println("\n")
	fmt.Printf("Temperature: %f", data.Current.TempC)
} 

func main() {
	fmt.Println("HELLO!")

	var r RequestData
//	r.city = "Minsk"
	r.air = "no"

	a := flag.String("city", "Minsk", "Enter City name for weather")
	flag.Parse()
	r.city = *a
	fmt.Println(r.city)
	
	res := build_url(r)
	content, _ := http.Get(res)
	defer content.Body.Close()
	body, _ := io.ReadAll(content.Body)
	bodystr := string(body)
	//fmt.Println(bodystr)

	parse_request(bodystr)
	fmt.Println("\n")
}
