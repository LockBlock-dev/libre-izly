package lib

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/LockBlock-dev/libre-izly/core"
	"github.com/LockBlock-dev/libre-izly/internal"
)

func GenerateQRCodeDataWithTimeAndPersistedAuthentification(
	mode core.QRCodeMode,
	datetime string,
	qrCodeVersion core.QRCodeVersion,
) (string, error) {
	auth, err := RetrieveAuthData()
	if err != nil {
		return "", err
	}

	qr, err := GenerateQRCodeDataWithTime(
		mode,
		auth.UserPublicId,
		datetime,
		qrCodeVersion,
		auth.NSSE,
		auth.Seed,
		auth.Counter,
		auth.QrCodePrivateKey,
	)
	if err != nil {
		return "", err
	}

	auth.Counter += 1

	PersistAuthData(auth)

	return qr, nil
}

func GenerateQRCodeDataWithPersistedAuthentification(
	mode core.QRCodeMode,
	qrCodeVersion core.QRCodeVersion,
) (string, error) {
	auth, err := RetrieveAuthData()
	if err != nil {
		return "", err
	}

	qr, err := GenerateQRCodeData(
		mode,
		auth.UserPublicId,
		qrCodeVersion,
		auth.NSSE,
		auth.Seed,
		auth.Counter,
		auth.QrCodePrivateKey,
	)
	if err != nil {
		return "", err
	}

	auth.Counter += 1

	PersistAuthData(auth)

	return qr, nil
}

func GenerateQRCodeData(
	mode core.QRCodeMode,
	userPublicId string,
	qrCodeVersion core.QRCodeVersion,
	nsse string,
	hotpSecret string,
	hotpCounter int,
	privateKey string,
) (string, error) {
	return GenerateQRCodeDataWithTime(
		mode,
		userPublicId,
		time.Now().UTC().Format(time.DateTime),
		qrCodeVersion,
		nsse,
		hotpSecret,
		hotpCounter,
		privateKey,
	)
}

func GenerateQRCodeDataWithTime(
	mode core.QRCodeMode,
	userPublicId string,
	datetime string,
	qrCodeVersion core.QRCodeVersion,
	nsse string,
	hotpSecret string,
	hotpCounter int,
	privateKey string,
) (string, error) {
	qr := fmt.Sprintf(
		"%s;%s;%s;%d",
		mode,
		userPublicId,
		datetime,
		qrCodeVersion,
	)

	hotp, err := internal.ComputeHOTPFromString(hotpSecret, hotpCounter)
	if err != nil {
		return "", err
	}

	log.Println(hotp)

	hash, err := internal.ComputeHMACHashFromStrings(fmt.Sprintf("%s+%s", qr, nsse), hotp)
	if err != nil {
		return "", err
	}

	qr += fmt.Sprintf(";%s;", hex.EncodeToString(hash))

	signature, err := internal.ComputeECDSASignatureFromStrings(qr, privateKey)
	if err != nil {
		return "", err
	}

	qr += base64.StdEncoding.EncodeToString(signature)

	return qr, nil
}
