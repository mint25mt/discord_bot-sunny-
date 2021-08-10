package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Weather struct {
	Area     string `json:"targetArea"`
	HeadLine string `json:"headlineText"`
	Body     string `json:"text"`
}

func GetWeather() string {
	jsonStr := httpGetStr("https://www.jma.go.jp/bosai/forecast/data/overview_forecast/230000.json")
	weather := formatWeather(jsonStr)

	area := fmt.Sprintf("%sの天気です。\n", weather.Area)
	head := fmt.Sprintf("%s\n", weather.HeadLine)
	body := fmt.Sprintf("%s\n", weather.Body)
	result := area + head + body

	return result
}

func httpGetStr(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Get Http Error:", err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("IO Read Error:", err)
	}
	defer response.Body.Close()
	return string(body)
}

func formatWeather(str string) *Weather {
	weather := new(Weather)
	if err := json.Unmarshal([]byte(str), weather); err != nil {
		log.Fatal("JSON Unmarshal error:", err)
	}
	return weather
}
