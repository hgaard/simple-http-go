package logger

import "os"

// https://stackoverflow.com/a/40326580
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
