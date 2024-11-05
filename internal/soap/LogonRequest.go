package soap

import (
	"encoding/xml"
)

type TypeElement struct {
	Type    string `xml:"i:type,attr"`
	Content string `xml:",chardata"`
}

type NullElement struct {
	Null string `xml:"i:null,attr"`
}

type LogonRequestEnvelope struct {
	XMLName xml.Name           `xml:"v:Envelope"`
	XMLNSI  string             `xml:"xmlns:i,attr"`
	XMLNSD  string             `xml:"xmlns:d,attr"`
	XMLNSC  string             `xml:"xmlns:c,attr"`
	XMLNSV  string             `xml:"xmlns:v,attr"`
	Header  LogonRequestHeader `xml:"v:Header"`
	Body    LogonRequestBody   `xml:"v:Body"`
}

type LogonRequestHeader struct{}

type LogonRequestBody struct {
	Logon Logon `xml:"Logon"`
}

type Logon struct {
	XMLName          xml.Name     `xml:"Logon"`
	XMLNS            string       `xml:"xmlns,attr"`
	ID               string       `xml:"id,attr"`
	Root             string       `xml:"c:root,attr"`
	Version          TypeElement  `xml:"version"`
	Channel          TypeElement  `xml:"channel"`
	Format           TypeElement  `xml:"format"`
	Model            TypeElement  `xml:"model"`
	Language         TypeElement  `xml:"language"`
	User             TypeElement  `xml:"user"`
	Password         interface{}  `xml:"password"`
	ActivationCode   *TypeElement `xml:"actCode,omitempty"`
	SmoneyClientType TypeElement  `xml:"smoneyClientType"`
	Rooted           TypeElement  `xml:"rooted"`
}

type LogonParams struct {
	Version          string
	Channel          string
	Format           string
	Model            string
	Language         string
	User             string
	Password         *string
	ActivationCode   *string
	SmoneyClientType string
	Rooted           string
}

func NewLogonRequestPayload(params *LogonParams) (string, error) {
	logon := Logon{
		XMLNS:            "Service",
		ID:               "o0",
		Root:             "1",
		Version:          TypeElement{Type: "d:string", Content: params.Version},
		Channel:          TypeElement{Type: "d:string", Content: params.Channel},
		Format:           TypeElement{Type: "d:string", Content: params.Format},
		Model:            TypeElement{Type: "d:string", Content: params.Model},
		Language:         TypeElement{Type: "d:string", Content: params.Language},
		User:             TypeElement{Type: "d:string", Content: params.User},
		SmoneyClientType: TypeElement{Type: "d:string", Content: params.SmoneyClientType},
		Rooted:           TypeElement{Type: "d:string", Content: params.Rooted},
	}

	if params.Password == nil && params.ActivationCode != nil {
		logon.Password = NullElement{Null: "true"}
		logon.ActivationCode = &TypeElement{
			Type:    "d:string",
			Content: *params.ActivationCode,
		}
	} else if params.Password != nil {
		logon.Password = TypeElement{
			Type:    "d:string",
			Content: *params.Password,
		}
	}

	envelope := LogonRequestEnvelope{
		XMLNSI: "http://www.w3.org/2001/XMLSchema-instance",
		XMLNSD: "http://www.w3.org/2001/XMLSchema",
		XMLNSC: "http://schemas.xmlsoap.org/soap/encoding/",
		XMLNSV: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: LogonRequestHeader{},
		Body: LogonRequestBody{
			Logon: logon,
		},
	}

	xmlData, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		return "", err
	}

	return xml.Header + string(xmlData), nil
}
