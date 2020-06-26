package utils

import (
	"strconv"
	"time"
)

// GenerateToken returns unique token for provided name
func GenerateToken(name string) string {
	return GetHash(name + "-" + strconv.FormatInt(time.Now().Unix(), 10) + "-" + NewRandomIntString(10))
}
