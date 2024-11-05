package soap

import (
	"encoding/xml"
	"errors"
)

type LogonResponseEnvelope struct {
	XMLName xml.Name          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    LogonResponseBody `xml:"Body"`
}

type LogonResponseBody struct {
	LogonResponse LogonResponse `xml:"LogonResponse"`
}

type LogonResponse struct {
	XMLName     xml.Name `xml:"LogonResponse"`
	XMLNS       string   `xml:"xmlns,attr"`
	LogonResult string   `xml:"LogonResult"`
}

type LogonError struct {
	XMLName xml.Name `xml:"E"`
	Error   string   `xml:"Error"`
	Code    string   `xml:"Code"`
	Msg     string   `xml:"Msg"`
	Title   string   `xml:"Title"`
}

type LogonResult struct {
	XMLName xml.Name `xml:"UserData"`
	UID     string   `xml:"UID"`
	Salt    string   `xml:"SALT"`
}

func ParseLogonResponse(xmlData string) (*LogonResult, error) {
	var envelope LogonResponseEnvelope

	err := xml.Unmarshal([]byte(xmlData), &envelope)
	if err != nil {
		return nil, err
	}

	var logonResult LogonResult
	err = xml.Unmarshal([]byte(envelope.Body.LogonResponse.LogonResult), &logonResult)
	if err != nil {
		return nil, err
	}

	var logonError LogonError
	err = xml.Unmarshal([]byte(envelope.Body.LogonResponse.LogonResult), &logonError)
	if err == nil {
		return nil, errors.New(logonError.Msg)
	}

	return &logonResult, nil
}

type ServicesInfos struct {
	XMLName       xml.Name `xml:"SERVICES_INFOS"`
	ServicesInfos []ServicesInfo
}

type ServicesInfo struct {
	XMLName    xml.Name `xml:"SERVICE"`
	Id         string   `xml:"ID"`
	CguExpired string   `xml:"CGU_EXPIRED"`
}

type Banks struct {
	XMLName xml.Name `xml:"BANKS"`
	Banks   []Bank   `xml:"BankCode"`
}

type Bank struct {
	Code string `xml:",chardata"`
}

type OAuth struct {
	XMLName      xml.Name `xml:"OAUTH"`
	AccessToken  string   `xml:"ACCESS_TOKEN"`
	TokenType    string   `xml:"TOKEN_TYPE"`
	ExpiresIn    string   `xml:"EXPIRES_IN"`
	RefreshToken string   `xml:"REFRESH_TOKEN"`
}

type Up struct {
	XMLName     xml.Name `xml:"UP"`
	Balance     string   `xml:"BAL"`
	CashBalance string   `xml:"CASHBAL"`
	LUD         string   `xml:"LUD"`
}

type LogonSecondStepResult struct {
	XMLName                         xml.Name `xml:"Logon"`
	Banks                           Banks
	UserId                          string `xml:"USER_ID"`
	Age                             string `xml:"AGE"`
	ZipCode                         string `xml:"ZIP_CODE"`
	Crous                           string `xml:"CROUS"`
	CategoryUserId                  string `xml:"CATEGORY_USERID"`
	Optin                           string `xml:"OPTIN"`
	ServicesInfos                   ServicesInfos
	UID                             string `xml:"UID"`
	SID                             string `xml:"SID"`
	P2PPAYMIN                       string `xml:"P2PPAYMIN"`
	P2PPAYMAX                       string `xml:"P2PPAYMAX"`
	P2PPAYPARTMIN                   string `xml:"P2PPAYPARTMIN"`
	P2PPAYPARTMAX                   string `xml:"P2PPAYPARTMAX"`
	Currency                        string `xml:"CUR"`
	MONEYINMIN                      string `xml:"MONEYINMIN"`
	MONEYINMAX                      string `xml:"MONEYINMAX"`
	MONEYOUTMIN                     string `xml:"MONEYOUTMIN"`
	MONEYOUTMAX                     string `xml:"MONEYOUTMAX"`
	NBP2PREQUEST                    string `xml:"NBP2PREQUEST"`
	NBP2PGET                        string `xml:"NBP2PGET"`
	Token                           string `xml:"TOKEN"`
	UserStatus                      string `xml:"USERSTATUS"`
	PROCASHIER                      string `xml:"PROCASHIER"`
	FirstName                       string `xml:"FNAME"`
	LastName                        string `xml:"LNAME"`
	Email                           string `xml:"EMAIL"`
	Alias                           string `xml:"ALIAS"`
	Status                          string `xml:"STATUS"`
	OPTINPARTNERS                   string `xml:"OPTINPARTNERS"`
	Blocked                         string `xml:"BLOCKED"`
	NewVersion                      string `xml:"NEWVERSION"`
	OAuth                           OAuth
	CguExpired                      string `xml:"CGU_EXPIRED"`
	PROG7                           string `xml:"PROG7"`
	CrousName                       string `xml:"CROUS_NAME"`
	SubscriptionDate                string `xml:"SUBSCRIPTION_DATE"`
	TermsAndConditionsAgreementDate string `xml:"TERMS_CONDITIONS_AGREEMENT_DATE"`
	Up                              Up
	HasNewActu                      string `xml:"HAS_NEW_ACTU"`
	Seed                            string `xml:"SEED"`
	NSSE                            string `xml:"NSSE"`
	UserPublicId                    string `xml:"USER_PUBLIC_ID"`
	QrCodePrivateKey                string `xml:"QR_CODE_PRIVATE_KEY"`
}

func ParseLogonSecondStepResponse(xmlData string) (*LogonSecondStepResult, error) {
	var envelope LogonResponseEnvelope

	err := xml.Unmarshal([]byte(xmlData), &envelope)
	if err != nil {
		return nil, err
	}

	var logonResult LogonSecondStepResult

	err = xml.Unmarshal([]byte(envelope.Body.LogonResponse.LogonResult), &logonResult)
	if err != nil {
		return nil, err
	}

	return &logonResult, nil
}
