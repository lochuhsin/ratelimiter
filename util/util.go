package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetLogger(packageLocation string) *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2021-01-01 11:11:11",
	})
	log.WithFields(logrus.Fields{
		"PackageLocation": packageLocation,
	})
	return log
}

type Settings struct {
	RedisPort       string
	PostgresPort    string
	PredefineString string
	Port            string
}

func GetSettings() Settings {
	return Settings{
		Port:            os.Getenv("PORT"),
		RedisPort:       os.Getenv("REDIS_PORT"),
		PostgresPort:    os.Getenv("POSTGRES_PORT"),
		PredefineString: os.Getenv("PREDEFINE_STRING"),
	}
}
