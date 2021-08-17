package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"coffeeBreak.com/m/v2/bot"
	"coffeeBreak.com/m/v2/corona"
	"coffeeBreak.com/m/v2/types"
	"coffeeBreak.com/m/v2/weather"
)

// Variables used for command line parameters
func main() {
	go bot.CatBot()
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		speakToDiscord(res, req)
	})
	// http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	http.ListenAndServe(":8080", nil)
}

func speakToDiscord(res http.ResponseWriter, req *http.Request) {
	value := "invalid parameter"
	switch req.FormValue("content") {
	case "corona":
		corona.DownloadFile("corona.csv", "https://covid19.mhlw.go.jp/public/opendata/newly_confirmed_cases_daily.csv")
		value = corona.DisPlayTodayCorona("corona.csv", req.FormValue("date"))
	case "weather":
		value = weather.GetWeather()
	}
	io.WriteString(res, value)
	// var username = os.Getenv("HOOKS_NAME")
	var username = "Captain Hook"
	var content = value
	// var url = os.Getenv("URL")
	var url = "https://discordapp.com/api/webhooks/874657070807388200/nhYHdQ9up7IwTg2GnHy3Gdnfm1Ki7HqDcYANN6FlOz-WHrp0C3mOz5k5ijkJe9J62qZU"

	message := types.Message{
		Username: &username,
		Content:  &content,
	}

	err := SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}

}

func SendMessage(url string, message types.Message) error {
	payload := new(bytes.Buffer)

	err := json.NewEncoder(payload).Encode(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		return err
	}

	fmt.Println(resp.StatusCode)
	if !((resp.StatusCode >= 200) && (resp.StatusCode < 300)) {
		defer resp.Body.Close()
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf(string(responseBody))
	}

	return nil
}
