package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port                = 0
	DateLayout          = "2006-01-02"
	AbsolutePath        = ""
	UAParserRegexesPath = ""
	SentryDSN           = ""
	DiscordUrl          = ""
)

func Load() {
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))

	if err != nil || debug {
		debug = true

		err = godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	_, b, _, _ := runtime.Caller(0)
	AbsolutePath = filepath.Join(filepath.Dir(b), "../..")

	UAParserRegexesPath = fmt.Sprintf("%s/regexes.yaml", AbsolutePath)

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 8000
	}

	DiscordUrl = os.Getenv("DISCORD_URL")

}
