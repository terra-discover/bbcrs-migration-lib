package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporatePayload struct
type CorporatePayload struct {
	basic.Base
	basic.DataOwner
	CorporatePayloadAPI
	Corporate *Corporate `json:"corporate,omitempty"`
}

// CorporatePayloadAPI for data payload
type CorporatePayloadAPI struct {
	CorporateID   *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Corporate ID
	PayloadStatus *int       `json:"payload_status,omitempty" example:"432"`                                                     // Payload Status
	Request       *string    `json:"request,omitempty" gorm:"type:varchar(4000)"`                                                // Request
	Response      *string    `json:"response,omitempty" gorm:"type:varchar(4000)"`                                               // Response
}
