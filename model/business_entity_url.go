package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BusinessEntityURL Operation Schedule
type BusinessEntityURL struct {
	basic.Base
	basic.DataOwner
	BusinessEntityURLAPI
}

// BusinessEntityURLAPI Operation Schedule API
type BusinessEntityURLAPI struct {
	BusinessEntityID *uuid.UUID `json:"business_entity_id,omitempty" gorm:"type:varchar(36);not null;index:idx_business_entity_url__business_entity_id,unique,where:deleted_at is null;" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml // Operation Schedule Name
	URLTypeID        *uuid.UUID `json:"url_type_id,omitempty" gorm:"type:varchar(36);index:idx_business_entity_url__business_entity_id,unique,where:deleted_at is null;" swaggertype:"string" format:"uuid"`                 // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml // Operation Schedule Name
	URL              *string    `json:"url,omitempty" gorm:"type:varchar(256)" example:"https://google.com"`                                                                                                                 // URL
	IsPrimary        *bool      `json:"is_primary,omitempty" gorm:"type:bool;index:idx_business_entity_url__business_entity_id,unique,where:deleted_at is null;" example:"true"`                                             // Is Primary
}
