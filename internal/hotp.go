package internal

import (
	"encoding/base64"
)

func ComputeHOTP(secret []byte, counter int) (string, error) {
	counterBytes := make([]byte, 8)

	for i := 7; i >= 0; i-- {
		counterBytes[i] = byte(counter & 0xFF)
		counter >>= 8
	}

	hmac, err := ComputeHMACHash(counterBytes, secret)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(hmac), nil
}

func ComputeHOTPFromString(secret string, counter int) (string, error) {
	secretBytes, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	return ComputeHOTP(secretBytes, counter)
}
