package lib

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/LockBlock-dev/libre-izly/core"
	"github.com/LockBlock-dev/libre-izly/internal/soap"
)

const MIMETextXml string = "text/xml;charset=utf-8"
const UserAgent string = "ksoap2-android/2.6.0+"
const ACTIVATION_URL string = "https://mon-espace.izly.fr/tools/Activation/"

func NewLogonParams(user string, password string, language string) *soap.LogonParams {
	return &soap.LogonParams{
		Version:          os.Getenv("CLIENT_VERSION"),
		Channel:          "AIZ",
		Format:           "T",
		Model:            "A",
		Language:         language,
		User:             user,
		Password:         &password,
		SmoneyClientType: core.SMONEY_CLIENT_TYPE_PART,
		Rooted:           "0",
	}
}

func NewLogonSecondStepParams(user string, activationCode string, language string) *soap.LogonParams {
	return &soap.LogonParams{
		Version:          os.Getenv("CLIENT_VERSION"),
		Channel:          "AIZ",
		Format:           "T",
		Model:            "A",
		Language:         language,
		User:             user,
		ActivationCode:   &activationCode,
		SmoneyClientType: core.SMONEY_CLIENT_TYPE_PART,
		Rooted:           "0",
	}
}

type SoapClient struct {
	Http *http.Client
	Url  string
}

func NewSoapClient() *SoapClient {
	return &SoapClient{
		Http: &http.Client{
			Timeout: 1500 * time.Millisecond,
		},
		Url: os.Getenv("SOAP_API_URL"),
	}
}

func (c *SoapClient) Logon(params *soap.LogonParams) (*soap.LogonResult, error) {
	payload, err := soap.NewLogonRequestPayload(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, c.Url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", MIMETextXml)
	request.Header.Add("User-Agent", UserAgent)
	request.Header.Add("smoneyClientType", params.SmoneyClientType)
	request.Header.Add("clientVersion", os.Getenv("CLIENT_VERSION"))

	resp, err := c.Http.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(b))

	return soap.ParseLogonResponse(string(b))
}

func (c *SoapClient) LogonSecondStep(params *soap.LogonParams) (*soap.LogonSecondStepResult, error) {
	payload, err := soap.NewLogonRequestPayload(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, c.Url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", MIMETextXml)
	request.Header.Add("User-Agent", UserAgent)
	request.Header.Add("smoneyClientType", params.SmoneyClientType)
	request.Header.Add("clientVersion", os.Getenv("CLIENT_VERSION"))

	resp, err := c.Http.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data, err := soap.ParseLogonSecondStepResponse(string(b))
	if err != nil {
		return nil, err
	}

	PersistAuthData(&AuthentificationData{
		UserId:           data.UserId,
		Seed:             data.Seed,
		Counter:          0,
		NSSE:             data.NSSE,
		UserPublicId:     data.UserPublicId,
		QrCodePrivateKey: data.QrCodePrivateKey,
	})

	return data, nil
}

func FetchActivationCode(userId string, activationId string) (string, error) {

	client := http.Client{
		Timeout: 1500 * time.Millisecond,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", ACTIVATION_URL, userId, activationId), nil)
	if err != nil {
		return "", err
	}

	request.Header.Add("User-Agent", os.Getenv("MOBILE_USER_AGENT"))

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	location := resp.Header.Get("Location")

	re := regexp.MustCompile(`izly:\/\/SBSCR\/\d{11}\/(?P<code>[\d\w]+)`)

	matches := re.FindStringSubmatch(location)
	lastIndex := re.SubexpIndex("code")

	if lastIndex == -1 {
		return "", errors.New("could not find activation code in Location header")
	}

	return matches[lastIndex], nil
}
