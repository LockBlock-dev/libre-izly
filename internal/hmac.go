package internal

import (
	"crypto/hmac"
	"crypto/sha1"
)

func ComputeHMACHash(message []byte, key []byte) ([]byte, error) {
	mac := hmac.New(sha1.New, key)

	_, err := mac.Write(message)
	if err != nil {
		return nil, err
	}

	return mac.Sum(nil), nil
}

func ComputeHMACHashFromStrings(message string, key string) ([]byte, error) {
	return ComputeHMACHash([]byte(message), []byte(key))
}
