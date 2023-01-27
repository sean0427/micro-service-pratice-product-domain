package config

import (
	"errors"
	"os"
)

func GetMongoDBUser() (string, error) {
	if v, found := os.LookupEnv("MONGO_APP_DB_USERNAME"); found {
		return v, nil
	}

	return "", errors.New("MONGO_APP_DB_USERNAME not found")
}

func GetMongoDBPassword() (string, error) {
	if v, found := os.LookupEnv("MONGO_APP_DB_PASSWORD"); found {
		return v, nil
	}

	return "", errors.New("MONGO_APP_DB_PASSWORD not found")
}

func GetMongoDBName() (string, error) {
	if v, found := os.LookupEnv("MONGODB_DB"); found {
		return v, nil
	}

	return "", errors.New("MONGODB_DB not found")
}

func GetMongoDBAddress() (string, error) {
	if v, found := os.LookupEnv("MONGODB_ADDRESS"); found {
		return v, nil
	}

	return "", errors.New("MONGODB_ADDRESS not found")
}

func GetMongoDBPort() string {
	if v, found := os.LookupEnv("MONGODB_PORT"); found {
		return v
	}

	return "27017"
}
