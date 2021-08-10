package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"coffeeBreak.com/m/v2/corona"
	"coffeeBreak.com/m/v2/weather"
)

type Message struct {
	Username  *string  `json:"username,omitempty"`
	AvatarUrl *string  `json:"avatar_url,omitempty"`
	Content   *string  `json:"content,omitempty"`
	Embeds    *[]Embed `json:"embeds,omitempty"`
}

type Embed struct {
	Title       *string    `json:"title,omitempty"`
	Url         *string    `json:"url,omitempty"`
	Description *string    `json:"description,omitempty"`
	Color       *string    `json:"color,omitempty"`
	Author      *Author    `json:"author,omitempty"`
	Fields      *[]Field   `json:"fields,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Image       *Image     `json:"image,omitempty"`
	Footer      *Footer    `json:"footer,omitempty"`
}

type Author struct {
	Name    *string `json:"name,omitempty"`
	Url     *string `json:"url,omitempty"`
	IconUrl *string `json:"icon_url,omitempty"`
}

type Field struct {
	Name   *string `json:"name,omitempty"`
	Value  *string `json:"value,omitempty"`
	Inline *bool   `json:"inline,omitempty"`
}

type Thumbnail struct {
	Url *string `json:"url,omitempty"`
}

type Image struct {
	Url *string `json:"url,omitempty"`
}

type Footer struct {
	Text    *string `json:"text,omitempty"`
	IconUrl *string `json:"icon_url,omitempty"`
}

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
		corona.DownloadFile("corona.csv", "https://covid19.mhlw.go.jp/public/opendata/newly_confirmed_cases_daily.csv")
		value = corona.DisPlayTodayCorona("corona.csv")
	case "weather":
		value = weather.GetWeather()
	}
	io.WriteString(res, value)
	var username = os.Getenv("HOOKS_NAME")
	var content = value
	var url = os.Getenv("URL")

	message := Message{
		Username: &username,
		Content:  &content,
	}

	err := SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("test")
}

func SendMessage(url string, message Message) error {
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
