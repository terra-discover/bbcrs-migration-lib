package model

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// FlightCachingSources
type FlightCachingSources struct {
	basic.Base
	basic.DataOwner
	basic.AuthContext
	SessionID            *uuid.UUID       `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_sources_session_id;not null"`
	ExpiredIn            *int64           `json:"expired_in,omitempty"`
	SourceType           *string          `json:"source_type,omitempty" gorm:"type:varchar(36);not null"`
	SourceTypeVisibility *string          `json:"source_type_visibility,omitempty" gorm:"type:varchar(200)"`
	ProgressStatus       *int             `json:"progress_status,omitempty" gorm:"type:smallint;default:0" example:"1" swaggerignore:"true"` // status (0: proposed, 1: reserved, 2:cancelled, 3:expired, 4:hold-confirmed, 5:confirmed)
	IsRetrieves          *bool            `json:"is_retrieves,omitempty"`                                                                    // IsNeedRepricing
	IsNeedRepricing      *bool            `json:"is_need_repricing,omitempty"`                                                               // IsRetrieves
	IsModifiedBooking    *bool            `json:"is_modified_booking,omitempty"`                                                             // This method is exclusively flag for case Modify flight after and before issued
	IsModifiedProduct    *bool            `json:"is_modified_product,omitempty"`                                                             // This method is exclusively flag for case Modify Seat Selection and Add-on Selection
	OldSessionID         *uuid.UUID       `json:"old_session_id,omitempty" gorm:"type:varchar(36)"`
	CreatedAt            *strfmt.DateTime `json:"created_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggerignore:"true"` // created at automatically inserted on post
	UpdatedAt            *strfmt.DateTime `json:"updated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggerignore:"true"` // updated at automatically changed on put or add on post
}

// UpdateProgressStatus
func (d *FlightCachingSources) UpdateProgressStatus(tx *gorm.DB, sessionID *uuid.UUID, value *int) {
	tx.Model(&d).Where(tx.Where(FlightCachingSources{
		SessionID: sessionID,
	})).Update("progress_status", value)
}

// BeforeCreate Data
func (d *FlightCachingSources) BeforeCreate(tx *gorm.DB) error {
	now := strfmt.DateTime(time.Now())
	d.CreatedAt = &now
	d.UpdatedAt = &now
	return nil
}

// BeforeUpdate Data
func (d *FlightCachingSources) BeforeUpdate(tx *gorm.DB) error {
	now := strfmt.DateTime(time.Now())
	d.UpdatedAt = &now
	return nil
}
