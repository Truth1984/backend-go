package util

import (
	"os"
)

func EnvHas(key string) bool {
	if v := os.Getenv(key); v != "" {
		return true
	}
	return false
}

func EnvGet(key string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return ""
}

func EnvGetExist(keys ...string) map[string]interface{} {
	m := make(map[string]interface{})
	for _, key := range keys {
		if EnvHas(key) {
			m[key] = EnvGet(key)
		}
	}
	return m
}

func EnvSet(key, value string) {
	os.Setenv(key, value)
}

func EnvSetOnce(key, value string) {
	if !EnvHas(key) {
		EnvSet(key, value)
	}
}
