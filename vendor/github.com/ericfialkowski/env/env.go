// Package env provides simple environment variable retrieval and parsing into builtin
// types including providing the default value if the environment variable is not set
// or not in the correct format.
package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// StringOrDefault returns the value in the system environment denoted by key or
// the supplied expectedValue if there is no environment variable named key.
func StringOrDefault(key, defaultValue string) string {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return envVal
}

// MustString returns the value in the system environment denoted by key or
// panics. Should really only be used when you can't run without the value and there
// is no viable default to provide
func MustString(key string) string {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("Missing required environment variable value for %s", key))
	}
	return envVal
}

// String returns the same as os.LookupEnv. This is just a simple
// wrapper for completeness with the other parsing methods.
func String(key string) (string, bool) {
	return os.LookupEnv(key)
}

// BoolOrDefault returns the value in the system environment denoted by key as
// a bool or the supplied expectedValue if there is no environment variable named key or
// if the value retrieved is not parsable as a bool.
func BoolOrDefault(key string, defaultValue bool) bool {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	r, err := strconv.ParseBool(envVal)
	if err != nil {
		return defaultValue
	}
	return r
}

// Bool returns the value in the system environment denoted by key as
// a bool and true. If there is no environment variable named key or
// if the value retrieved is not parsable as a bool then (false, false) is returned
func Bool(key string) (bool, bool) {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return false, false
	}
	r, err := strconv.ParseBool(envVal)
	if err != nil {
		return false, false
	}
	return r, true
}

// IntOrDefault returns the value in the system environment denoted by key as
// an int or the supplied expectedValue if there is no environment variable named key or
// if the value retrieved is not parsable as an int.
func IntOrDefault(key string, defaultValue int) int {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultValue
	}
	r, err := strconv.Atoi(envVal)
	if err != nil {
		return defaultValue
	}
	return r
}

// Int returns the value in the system environment denoted by key as
// an int and true. If there is no environment variable named key or
// if the value retrieved is not parsable as an int then (0, false) is returned
func Int(key string) (int, bool) {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return 0, false
	}
	r, err := strconv.Atoi(envVal)
	if err != nil {
		return 0, false
	}
	return r, true
}

// DurationOrDefault returns the value in the system environment denoted by key as
// a time.Duration or the supplied expectedValue if there is no environment variable named key or
// if the value retrieved is not parsable as a time.Duration. See time.ParseDuration() for the
// allowed formatting of the environment variable.
func DurationOrDefault(key string, defaultValue time.Duration) time.Duration {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultValue
	}
	r, err := time.ParseDuration(envVal)
	if err != nil {
		return defaultValue
	}
	return r
}

// Duration returns the value in the system environment denoted by key as
// a time.Duration and true. If there is no environment variable named key or
// if the value retrieved is not parsable as a time.Duration then (0, false) is returned
// if the value retrieved is not parsable as a time.Duration then (0, false) is returned
// See time.ParseDuration() for the allowed formatting of the environment variable.
func Duration(key string) (time.Duration, bool) {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return 0, false
	}
	r, err := time.ParseDuration(envVal)
	if err != nil {
		return 0, false
	}
	return r, true
}

// Float32OrDefault returns the value in the system environment denoted by key as
// an int or the supplied expectedValue if there is no environment variable named key or
// if the value retrieved is not parsable as a float32.
func Float32OrDefault(key string, defaultValue float32) float32 {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultValue
	}
	r, err := strconv.ParseFloat(envVal, 64)
	if err != nil {
		return defaultValue
	}
	return float32(r)
}

// Float32 returns the value in the system environment denoted by key as
// an int and true. If there is no environment variable named key or
// if the value retrieved is not parsable as a float32 then (0, false) is returned
func Float32(key string) (float32, bool) {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return 0, false
	}
	r, err := strconv.ParseFloat(envVal, 64)
	if err != nil {
		return 0, false
	}
	return float32(r), true
}

// Float64OrDefault returns the value in the system environment denoted by key as
// an int or the supplied expectedValue if there is no environment variable named key or
// if the value retrieved is not parsable as a float64.
func Float64OrDefault(key string, defaultValue float64) float64 {
	envVal := os.Getenv(key)
	if envVal == "" {
		return defaultValue
	}
	r, err := strconv.ParseFloat(envVal, 64)
	if err != nil {
		return defaultValue
	}
	return r
}

// Float64 returns the value in the system environment denoted by key as
// an int and true. If there is no environment variable named key or
// if the value retrieved is not parsable as a float64 then (0, false) is returned
func Float64(key string) (float64, bool) {
	envVal, ok := os.LookupEnv(key)
	if !ok {
		return 0, false
	}
	r, err := strconv.ParseFloat(envVal, 64)
	if err != nil {
		return 0, false
	}
	return r, true
}
