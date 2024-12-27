package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BusinessEntityEmail Operation Schedule
type BusinessEntityEmail struct {
	basic.Base
	basic.DataOwner
	BusinessEntityEmailAPI
	BusinessEntity   *BusinessEntity   `json:"business_entity,omitempty"`
	EmailAddressType *EmailAddressType `json:"email_address_type,omitempty"`
}

// BusinessEntityEmailAPI Operation Schedule API
type BusinessEntityEmailAPI struct {
	BusinessEntityID   *uuid.UUID `json:"business_entity_id,omitempty" gorm:"type:varchar(36);not null;index:idx_business_entity_email__business_entity_id,unique,where:deleted_at is null;" swaggertype:"string" format:"uuid"`       // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml // Operation Schedule Name
	EmailAddressTypeID *uuid.UUID `json:"email_address_type_id,omitempty" gorm:"type:varchar(36);index:idx_business_entity_email__business_entity_id,unique,where:deleted_at is null;" swaggertype:"string" format:"uuid"`             // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml // Operation Schedule Name
	Email              *string    `json:"email,omitempty" gorm:"type:varchar(256);check:chk_business_entity_email__email_lowercase,coalesce(email=lower(email),true)=true;" example:"anonymous@nomail.com" validate:"omitempty,email"` // Email
	IsPrimary          *bool      `json:"is_primary,omitempty"  gorm:"default:false;" example:"true"`                                                                                                                                  // Is Primary
}
