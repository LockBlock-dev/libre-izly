package internal

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
)

func ComputeECDSASignature(message []byte, privateKey []byte) ([]byte, error) {
	key, err := x509.ParsePKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(message)

	signature, err := ecdsa.SignASN1(rand.Reader, key.(*ecdsa.PrivateKey), hash[:])
	if err != nil {
		return nil, err
	}

	return signature, nil
}

func ComputeECDSASignatureFromStrings(message string, privateKey string) ([]byte, error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	return ComputeECDSASignature([]byte(message), privateKeyBytes)
}
