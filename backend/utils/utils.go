package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func InitializeLogging() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func ConfigureGinMod() {
	env := os.Getenv("ENVIRONMENT")

	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func GetPortNumber() int {
	port := os.Getenv("PORT")

	if port == "" {
		log.Warn().Msg("Port environment variable was not provided using default port 5000")
		return 5000
	}

	portNumber, err := strconv.Atoi(port)

	if err != nil {
		log.Fatal().Msgf("Failed to parse port number: %s", err)
	}

	return portNumber
}

func SqlLike(value string) string {
	return "%" + value + "%"
}
func GetEnvironmentVariable(key string) string {
	var environment = os.Getenv(key)

	if environment == "" {
		log.Fatal().
			Msgf("Failed to find environment variable \"%s\"", key)

		panic("Missing Environment variable")
	}

	return environment
}
