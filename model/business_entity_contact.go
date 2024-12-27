package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BusinessEntityContact struct
type BusinessEntityContact struct {
	basic.Base
	basic.DataOwner
	BusinessEntityContactAPI
	BusinessEntity   *BusinessEntity `json:"business_entity,omitempty"`                                                                                                                            // BusinessEntity
	PersonID         *uuid.UUID      `json:"person_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                                       // Person ID
	Person           *Person         `json:"person,omitempty"`                                                                                                                                     // Person
	BusinessEntityID *uuid.UUID      `json:"business_entity_id,omitempty" gorm:"type:varchar(36);not null;index:idx_business_entity_contact__business_entity_id,unique,where:deleted_at is null;"` // Business Entity Id
	ContactTypeID    *uuid.UUID      `json:"contact_type_id,omitempty" gorm:"type:varchar(36);index:idx_business_entity_contact__business_entity_id,unique,where:deleted_at is null;"`             // Contact Type Id
}

// BusinessEntityContactAPI for secure request body API
type BusinessEntityContactAPI struct {
	ContactName        *string `json:"contact_name,omitempty" gorm:"type:varchar(128)" example:"John Doe"`                                                                                                                                                     // Contact Name
	ContactPhoneNumber *string `json:"contact_phone_number,omitempty" gorm:"type:varchar(32)" example:"+62812345678910"`                                                                                                                                       // Contact Phone Number
	ContactEmail       *string `json:"contact_email,omitempty" gorm:"type:varchar(256);check:chk_business_entity_contact__contact_email_lowercase,coalesce(contact_email=lower(contact_email),true)=true;" example:"demo@mail.com" validate:"omitempty,email"` // Contact Email
	Relationship       *string `json:"relationship,omitempty" gorm:"type:varchar(36)" example:"Neighbor"`                                                                                                                                                      // Relationship
	IsPrimary          *bool   `json:"is_primary,omitempty" gorm:"type:bool;index:idx_business_entity_contact__business_entity_id,unique,where:deleted_at is null;"`                                                                                           // Is Primary
}
