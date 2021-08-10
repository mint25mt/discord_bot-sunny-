package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"coffeeBreak.com/m/v2/weather"
	"github.com/gtuk/discordwebhook"
)

// Variables used for command line parameters
func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		speakToDiscord(res, req)
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func speakToDiscord(res http.ResponseWriter, req *http.Request) {
	value := "invalid parameter"
	switch req.FormValue("content") {
	case "corona":
		// corona.DownloadFile("corona.csv", "https://covid19.mhlw.go.jp/public/opendata/newly_confirmed_cases_daily.csv")
		// value = corona.DisPlayTodayCorona("corona.csv")
		value = "corona"
	case "weather":
		value = weather.GetWeather()
	default:
		io.WriteString(res, value)
	}
	var username = os.Getenv("HOOKS_NAME")
	var content = value
	var url = os.Getenv("URL")

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
}
