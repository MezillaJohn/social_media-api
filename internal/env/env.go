package env

import (
	"fmt"
	"os"
	"strconv"
)

func GetStringEnv(key, fallback string) string {
	val, ok := os.LookupEnv(key)

	fmt.Print(val)

	if !ok {
		return fallback
	}

	return val
}

func GetIntEnv(key string, fallback int) int {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	intVal, err := strconv.Atoi(val)

	if err != nil {
		return fallback
	}

	return intVal
}
