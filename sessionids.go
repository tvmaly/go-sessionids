package sessionids

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// GenerateSessionId returns a random string given a length
// if the length is less than 1 a default length of 40 is used
// if any error occurs an empty string will be returned
func GenerateSessionId(length int) string {

	if length < 1 {
		length = 40
	}

	randbytes := make([]byte, length)

	if _, err := io.ReadFull(rand.Reader, randbytes); err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(randbytes)

}
