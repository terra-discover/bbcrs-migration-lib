package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type SabreTraceLog struct {
	basic.Base
	basic.DataOwner
	SessionID      *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Action         *string    `json:"action,omitempty" gorm:"type:varchar(50)"`
	Payload        *string    `json:"payload,omitempty" gorm:"type:text"`
	Request        *string    `json:"request,omitempty" gorm:"type:text"`
	Response       *string    `json:"response,omitempty" gorm:"type:text"`
	ResponseStatus *string    `json:"resp_status,omitempty" gorm:"type:text"`
	ResponseWarn   *string    `json:"resp_warn,omitempty" gorm:"type:text"`
	ResponseError  *string    `json:"resp_error,omitempty" gorm:"type:text"`
	ResponseFault  *string    `json:"resp_fault,omitempty" gorm:"type:text"`
	ResponseXml    *string    `json:"response_xml,omitempty" gorm:"type:text"`
	Expired        *string    `json:"expired,omitempty"`
}

// TODO Collect Response Error (Response Code) from sabre
