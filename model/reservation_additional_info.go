package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReservationAdditionalInfo  Reservation Additional Info
type ReservationAdditionalInfo struct {
	basic.Base
	basic.DataOwner
	ReservationAdditionalInfoAPI
	ReservationAdditionalInfoSelection []ReservationAdditionalInfoSelection `json:"reservation_additional_info_selection,omitempty"`
	AdditionalInfo                     *AdditionalInfo                      `json:"additional_info,omitempty"`
}

// ReservationAdditionalInfoAPI Reservation Additional Info API
type ReservationAdditionalInfoAPI struct {
	ReservationID       *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	AdditionalInfoID    *uuid.UUID `json:"additional_info_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	AdditionalInfoName  *string    `json:"additional_info_name,omitempty" gorm:"type:varchar(256)"`
	AdditionalInfoValue *string    `json:"additional_info_value,omitempty" gorm:"text"`
}

type ReservationAdditionalInformationResponse struct {
	ReservationID         *uuid.UUID                  `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	SessionID             *uuid.UUID                  `json:"session_id,omitempty"  gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	AdditionalInformation []ReservationAdditionalInfo `json:"additional_information"`
	Project               *Project                    `json:"project"`
	CostCenter            *CostCenter                 `json:"cost_center"`
	TravelPurpose         *TravelPurpose              `json:"travel_purpose"`
}
