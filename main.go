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
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/drive/files"
	"github.com/yitsushi/go-misskey/services/notes"
)

type MastodonConfig struct {
	Enable       bool   `json:"enable"`
	Instance     string `json:"instance"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
}

type MisskeyConfig struct {
	Enable      bool   `json:"enable"`
	Instance    string `json:"instance"`
	AccessToken string `json:"token"`
}

type MisskeyImage struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	MD5         string    `json:"md5"`
	Size        int       `json:"size"`
	IsSensitive bool      `json:"isSensitive"`
	Blurhash    string    `json:"blurhash"`
	Properties  struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"properties"`
	URL          string      `json:"url"`
	ThumbnailURL string      `json:"thumbnailUrl"`
	Comment      interface{} `json:"comment"`
	FolderID     interface{} `json:"folderId"`
	Folder       interface{} `json:"folder"`
	UserID       interface{} `json:"userId"`
	User         interface{} `json:"user"`
}

type Config struct {
	Mastodon MastodonConfig `json:"mastodon"`
	Misskey  MisskeyConfig  `json:"misskey"`
	City     string         `json:"city"`
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
	fmt.Println(config.Misskey.AccessToken)

	// downloads the Image
	getWeatherData(config.City)

	if config.Misskey.Enable {
		misskeyPost(config)
	}
	if config.Mastodon.Enable {
		mastodonPost(config)
	}

	os.Remove("weather.png")
}

func misskeyPost(config Config) {
	client, err := misskey.NewClientWithOptions(
		misskey.WithAPIToken(config.Misskey.AccessToken),
		misskey.WithBaseURL("https", config.Misskey.Instance, ""),
	)

	// upload the image to the Misskey instance
	fileContent, err := os.ReadFile("weather.png")
	if err != nil {
		panic(err)
	}

	file, err := client.Drive().File().Create(files.CreateRequest{
		FolderID:    "9iw1wl10xce0ik2g",
		IsSensitive: false,
		Force:       false,
		Content:     fileContent,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(file.URL)
	fmt.Println(file.ID)

	resp, err := client.Notes().Create(notes.CreateRequest{
		Text:       core.NewString("Hier ist das heutige Wetter von " + config.City + ". Ich wünsche euch einen schönen " + getWeekDay() + "!"),
		Visibility: models.VisibilityFollowers,
		FileIDs: []string{
			file.ID,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.CreatedNote.URL)
}

func mastodonPost(config Config) {
	// Initializing the client
	client := mastodon.NewClient(&mastodon.Config{
		Server:       config.Mastodon.Instance,
		ClientID:     config.Mastodon.ClientID,
		ClientSecret: config.Mastodon.ClientSecret,
		AccessToken:  config.Mastodon.AccessToken,
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
	fmt.Println("Posted the toot successfully at " + config.Mastodon.Instance + ": " + toot.URL)
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
