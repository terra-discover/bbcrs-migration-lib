package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ReservationAdditionalInfoSelection struct {
	basic.Base
	basic.DataOwner
	ReservationAdditionalInfoSelectionAPI
	ReservationAdditionalInfo *ReservationAdditionalInfo `json:"reservation_additional_info,omitempty" gorm:"foreignKey:ReservationAdditionalInfoID;references:ID"`
}

type ReservationAdditionalInfoSelectionAPI struct {
	ReservationAdditionalInfoID  *uuid.UUID `json:"reservation_additional_info_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	AdditionalInfoOptionID       *uuid.UUID `json:"additional_info_option_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string"`
	AdditionalInfoOptionName     *string    `json:"additional_info_option_name,omitempty" gorm:"type:varchar(256)"`
	AdditionalInfoSelectionValue *string    `json:"additional_info_selection_value,omitempty" gorm:"type:text"`
}
