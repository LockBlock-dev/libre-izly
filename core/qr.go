package core

type QRCodeMode = string

const (
	QR_CODE_MODE_IZLY   QRCodeMode = "AIZ"
	QR_CODE_MODE_SMONEY QRCodeMode = "A"
)

type QRCodeVersion = int

const (
	QR_CODE_VERSION_TWO QRCodeVersion = iota + 2
	QR_CODE_VERSION_THREE
)
