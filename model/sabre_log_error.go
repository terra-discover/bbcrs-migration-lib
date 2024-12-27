package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type SabreLogError struct {
	basic.Base
	basic.DataOwner
	SessionID *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Type      *string    `json:"type,omitempty" gorm:"type:varchar(10)"` // success, warning, error
	Service   *string    `json:"service,omitempty" gorm:"type:varchar(60)"`
	Code      *string    `json:"code,omitempty" gorm:"type:varchar(4000)"`
	ShortText *string    `json:"short_text,omitempty" gorm:"type:varchar(4000)"`
	Message   *string    `json:"message,omitempty" gorm:"type:varchar(4000)"`
	Timestamp *string    `json:"timestamp,omitempty" gorm:"type:varchar(400)"`
}

// TODO: define structs below in integration-service model

// type SabreMapError struct {
// 	SabreErrorResponse
// 	SabreApplicationResults
// }

// type SabreFaultResponse struct {
// 	XMLName       xml.Name            `xml:"Envelope,omitempty"`
// 	SoapEnv       string              `xml:"soap-env,attr,omitempty"`
// 	SoapEnvHeader SabreResponseHeader `xml:"Header,omitempty"`
// 	SoapEnvBody   SabreFaultBody      `xml:"Body,omitempty"`
// }
// type SabreFaultBody struct {
// 	Fault struct {
// 		Text        string `xml:",chardata"`
// 		Faultcode   string `xml:"faultcode"`
// 		Faultstring string `xml:"faultstring"`
// 		Detail      struct {
// 			StackTrace string `xml:"StackTrace"`
// 		} `xml:"detail"`
// 	} `xml:"Fault"`
// }

// type SabreApplicationResults struct {
// 	Xmlns   string `xml:"xmlns,attr"`
// 	Status  string `xml:"status,attr"`
// 	Success struct {
// 		Text                  string                     `xml:",chardata"`
// 		TimeStamp             string                     `xml:"timeStamp,attr"`
// 		SystemSpecificResults SabreSystemSpecificResults `xml:"SystemSpecificResults"`
// 	} `xml:"Success"`
// 	Error []struct {
// 		Text                  string                     `xml:",chardata"`
// 		Type                  string                     `xml:"type,attr"`
// 		TimeStamp             string                     `xml:"timeStamp,attr"`
// 		SystemSpecificResults SabreSystemSpecificResults `xml:"SystemSpecificResults"`
// 	} `xml:"Error"`
// 	Warning []struct {
// 		Text                  string                     `xml:",chardata"`
// 		Type                  string                     `xml:"type,attr"`
// 		TimeStamp             string                     `xml:"timeStamp,attr"`
// 		SystemSpecificResults SabreSystemSpecificResults `xml:"SystemSpecificResults"`
// 	} `xml:"Warning"`
// }

// type SabreSystemSpecificResults struct {
// 	Text        string `xml:",chardata"`
// 	HostCommand struct {
// 		Text   string `xml:",chardata"`
// 		LNIATA string `xml:"LNIATA,attr"`
// 	} `xml:"HostCommand"`
// 	Message []struct {
// 		Text string `xml:",chardata"`
// 		Code string `xml:"code,attr"`
// 	} `xml:"Message"`
// 	ShortText string `xml:"ShortText"`
// }

// type SabreErrorResponse struct {
// 	Success  string
// 	Warnings struct {
// 		Warning []struct {
// 			Text         string `xml:",chardata"`
// 			Type         string `xml:"Type,attr"`
// 			Code         string `xml:"Code,attr"`
// 			MessageClass string `xml:"MessageClass,attr"`
// 			ShortText    string `xml:"ShortText,attr"`
// 		} `xml:"Warning"`
// 	} `xml:"Warnings"`
// 	Errors struct {
// 		Error []struct {
// 			Text      string `xml:",chardata"`
// 			Type      string `xml:"Type,attr"`
// 			Code      string `xml:"Code,attr"`
// 			ShortText string `xml:"ShortText,attr"`
// 		} `xml:"Error"`
// 	}
// }
