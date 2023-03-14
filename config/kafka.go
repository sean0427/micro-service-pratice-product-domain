package config

import (
	"errors"
	"os"
)

func GetKafkaPath() (string, error) {
	if v, found := os.LookupEnv("KAFKA_PATH"); found {
		return v, nil
	}

	return "", errors.New("KAFKA_PATH not found")
}
