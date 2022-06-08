package v1_utils

import (
	"math/rand"
	"time"
	"unsafe"
)

// How to generate a random string of a fixed length in Go?
// 8. "Mimicing" strings.Builder with package unsafe
// RandStringBytesMaskImprSrcUnsafe is 6.3 times faster than RandStringRunes(),
// uses one sixth memory and half as few allocations.
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

const (
	letterBytes   = "abcdefghkmnpqrstuvwxyzABCDEFGHKMNPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// RandomString generates a random string of size n
func RandomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	// Replace strings.Builder with unsafe
	return *(*string)(unsafe.Pointer(&b))
}
