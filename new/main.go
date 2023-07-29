package new

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
)

type Config struct {
	Token    string `json:"token"`
	City     string `json:"city"`
	Instance string `json:"instance"`
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

	client, err := misskey.NewClientWithOptions(
		misskey.WithAPIToken(config.Token),
		misskey.WithBaseURL("https", config.Instance, ""),
		misskey.WithLogLevel(logrus.DebugLevel),
	)
	if err != nil {
		logrus.Error(err.Error())
	}

	stats, err := client.Meta().Stats()
	if err != nil {
		log.Printf("[Meta] Error happened: %s", err)
		return
	}

	log.Printf("[Stats] Instances:          %d", stats.Instances)
	log.Printf("[Stats] NotesCount:         %d", stats.NotesCount)
	log.Printf("[Stats] UsersCount:         %d", stats.UsersCount)
	log.Printf("[Stats] OriginalNotesCount: %d", stats.OriginalNotesCount)
	log.Printf("[Stats] OriginalUsersCount: %d", stats.OriginalUsersCount)
}
