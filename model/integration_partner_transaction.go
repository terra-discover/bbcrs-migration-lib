package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerTransaction Integration Partner Transaction
type IntegrationPartnerTransaction struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerTransactionAPI
}

// IntegrationPartnerTransactionAPI Integration Partner Transaction API
type IntegrationPartnerTransactionAPI struct {
	IntegrationPartnerID        *uuid.UUID       `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null;"`                                    // Integration Partner ID
	IntegrationPartnerAirlineID *uuid.UUID       `json:"integration_partner_airline_id,omitempty" gorm:"type:varchar(36)"`                                      // Integration Partner Airline ID
	IntegrationPartnerHotelID   *uuid.UUID       `json:"integration_partner_hotel_id,omitempty" gorm:"type:varchar(36)"`                                        // Integration Partner Hotel ID
	TransactionType             *string          `json:"transaction_type,omitempty" gorm:"type:varchar(2)" example:"cc"`                                        // Transaction Type
	TransactionStatus           *string          `json:"transaction_status,omitempty" gorm:"type:smallint" example:"1"`                                         // Transaction Status
	TransactionTimestamp        *strfmt.DateTime `json:"transaction_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`       // Trasanction Timestamp
	TransactionLocalTimestamp   *strfmt.DateTime `json:"transaction_local_timestamp,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"` // Trasanction Local Timestamp
	Request                     *string          `json:"request,omitempty" gorm:"type:text"`                                                                    // Request
	Response                    *string          `json:"response,omitempty" gorm:"type:text"`                                                                   // Response
	Comment                     *string          `json:"comment,omitempty" gorm:"type:text"`                                                                    //Comment
}
