package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mattn/go-mastodon"
)

type Config struct {
	Instance     string `json:"instance"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
	City         string `json:"city"`
}

func main() {
	// check if config file exists
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		fmt.Println("config.json does not exist. Please create one using config.example.json as a template.")
		os.Exit(1)
	}

	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	client := mastodon.NewClient(&mastodon.Config{
		Server:       config.Instance,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		AccessToken:  config.AccessToken,
	})

	var ctx context.Context = context.TODO()

	// Show account information
	account, err := client.GetAccountCurrentUser(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var bot string
	if account.Bot {
		bot = "yes"
	} else {
		bot = "no"
	}

	fmt.Printf("Name: %s\nUsername: %s\nBot: %s\nFollowers: %d\n", account.DisplayName, account.Username, bot, account.FollowersCount)

	// Get weather data
	res, err := http.Get("https://wttr.in/" + config.City + "?format=%l:\\n+%" + "c+%C+%" + "t+%" + "w+%m\\n%20Pressure:+%P\\n%" + "20UV:+%u\\n\\n+%" + "T+%Z\\n")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	stringBody := string(body)
	message := fmt.Sprintf("Guten %s! Hier ist das Wetter für heute:\n\n%s", getWeekDay(), stringBody)

	// Post new status
	toot, err := client.PostStatus(ctx, &mastodon.Toot{
		Status:     message,
		Visibility: "unlisted",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\nToot!!!\n%s\n", toot.URL)
}

// ! PNG support is currently disabled on wttr.in
/**func getWeatherData(city string) {
	// Save image to file
	if _, err := os.Stat("weather.png"); err == nil {
		os.Remove("weather.png")
	}
	image, err := os.Create("weather.png")
	if err != nil {
		panic(err)
	}
	// retry every 30 seconds if the request fails
	c := http.Client{}
	var res *http.Response
	retry := true
	for retry {
		res, err = c.Get("https://wttr.in/" + city + ".png?1&lang=de")
		if err != nil {
			fmt.Println(err)
		} else {
			retry = false
		}
		res.Header.Set("User-Agent", "curl/7.54.0")
		defer res.Body.Close()
		time.Sleep(30 * time.Second)
	}
	io.Copy(image, res.Body)
	defer image.Close()
}**/

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
