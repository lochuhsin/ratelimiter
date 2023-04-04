package util

import "github.com/sirupsen/logrus"

func GetLogger(packageLocation string) *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2021-01-01 11:11:11",
		FieldMap: logrus.FieldMap{
			"PackageLocation": packageLocation,
		},
	})
	return log
}