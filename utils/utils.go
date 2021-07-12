package utils

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Str2int(str string, fallback int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return val
	}
	return fallback
}
