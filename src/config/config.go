package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port       = 0
	DateLayout = "2006-01-02"
	SentryDSN  = ""
	DiscordUrl = ""
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

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 8000
	}

	DiscordUrl = os.Getenv("DISCORD_URL")

}
