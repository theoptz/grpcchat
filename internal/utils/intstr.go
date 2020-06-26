package utils

import (
	"math/rand"
	"time"
	"unsafe"
)

const digits = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewRandomIntString returns string of specified length
// that contains from random digits
func NewRandomIntString(len int) string {
	if len < 1 {
		return ""
	}

	res := make([]byte, 0, len)

	for i := 0; i < len; i++ {
		res = append(res, digits[rand.Intn(10)])
	}

	return *(*string)(unsafe.Pointer(&res))
}
