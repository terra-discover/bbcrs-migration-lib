package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SiteContent SiteContent
type SiteContent struct {
	basic.Base
	basic.DataOwner
	SiteContentAPI
	Site               *Site               `json:"site,omitempty" gorm:"foreignKey:SiteID;references:ID"`
	ContentDescription *ContentDescription `json:"content_description,omitempty" gorm:"foreignKey:ContentDescriptionID;references:ID"`
}

// SiteContentAPI SiteContent API
type SiteContentAPI struct {
	SiteID               *uuid.UUID `json:"site_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	ContentDescriptionID *uuid.UUID `json:"content_description_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
}

// SiteContentResponse struct
type SiteContentResponse struct {
	SiteContent
	ContentDescriptionID     *uuid.UUID `json:"content_description_id,omitempty" gorm:"type:varchar(36)"`
	Title                    *string    `json:"title,omitempty" gorm:"type:varchar(256)"`
	URL                      *string    `json:"url,omitempty" gorm:"type:varchar(256)"`
	Description              *string    `json:"description,omitempty" gorm:"type:text"`
	InformationTypeID        *uuid.UUID `json:"information_type_id,omitempty" gorm:"type:varchar(36)"`
	InformationTypeCode      *int64     `json:"information_type_code,omitempty" gorm:"type:smallint;"`
	InformationTypeName      *string    `json:"information_type_name,omitempty" gorm:"type:varchar(256)"`
	AdditionalDetailTypeID   *uuid.UUID `json:"additional_detail_type_id,omitempty" gorm:"type:varchar(36)"`
	AdditionalDetailTypeCode *int       `json:"additional_detail_type_code,omitempty" gorm:"type:smallint;"`
	AdditionalDetailTypeName *string    `json:"additional_detail_type_name,omitempty" gorm:"type:varchar(256);"`
}
