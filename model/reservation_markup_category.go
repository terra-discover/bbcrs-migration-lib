package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ReservationMarkupCategory struct {
	basic.Base
	basic.DataOwner
	ReservationMarkupCategoryAPI
}

type ReservationMarkupCategoryAPI struct {
	ReservationID      *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	MarkupCategoryID   *uuid.UUID `json:"markup_category_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	MarkupCategoryCode *string    `json:"markup_category_code,omitempty" gorm:"type:varchar(36)"`
	MarkupCategoryName *string    `json:"markup_category_name,omitempty" gorm:"type:varchar(356)"`
	Description        *string    `json:"description,omitempty" gorm:"type:text"`
}
