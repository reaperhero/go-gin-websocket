package utils

import (
	"os"
	"strconv"
)

func GetEnvWithDefault(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}

func GetEnvInt(key string, defaultValue int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	ret, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}

	return ret
}
