package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/mattn/go-mastodon"
)

type Config struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
	City         string `json:"city"`
	Instance     string `json:"instance"`
}

func main() {
	// read the config.json file
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer configFile.Close()

	// parses the config file
	bytesOfConfig, _ := ioutil.ReadAll(configFile)
	var config Config
	json.Unmarshal(bytesOfConfig, &config)

	// downloads the Image
	getWeatherData(config.City)

	// Initializing the client
	client := mastodon.NewClient(&mastodon.Config{
		Server:       config.Instance,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		AccessToken:  config.AccessToken,
	})

	// uploading the image to the Mastodon instance
	file, err := os.ReadFile("weather.png")
	if err != nil {
		panic(err)
	}
	image, err := client.UploadMediaFromBytes(context.Background(), file)
	if err != nil {
		panic(err)
	}

	// creating a toot/status on the Mastodon instance
	toot, err := client.PostStatus(context.Background(), &mastodon.Toot{
		MediaIDs: []mastodon.ID{
			image.ID,
		},
		Language:   "de",
		Visibility: "unlisted",
		Status:     "Hier ist das heutige Wetter von " + config.City + ". Ich wünsche euch einen schönen " + getWeekDay() + "!",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Posted the toot successfully at " + config.Instance + ": " + toot.URL)
	os.Remove("weather.png")
}

func getWeatherData(city string) {
	// Save image to file
	if _, err := os.Stat("weather.png"); err == nil {
		os.Remove("weather.png")
	}
	image, err := os.Create("weather.png")
	if err != nil {
		panic(err)
	}
	c := http.Client{}
	res, err := c.Get("https://wttr.in/" + city + ".png?1&lang=de")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(image, res.Body)
	defer image.Close()
}

func getWeekDay() string {
	var day = time.Now().Weekday().String()
	switch day {
	case "Monday":
		return "Montag"
	case "Tuesday":
		return "Dienstag"
	case "Wednesday":
		return "Mittwoch"
	case "Thursday":
		return "Donnerstag"
	case "Friday":
		return "Freitag"
	case "Saturday":
		return "Samstag"
	case "Sunday":
		return "Sonntag"
	}
	return ""
}
