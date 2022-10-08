package mapper

import (
	"strings"
)

func MapKey(key string) string {
	splittedKey := strings.Split(key, "/")
	return splittedKey[len(splittedKey) - 1]
}
