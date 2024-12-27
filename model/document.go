package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Document struct
type Document struct {
	basic.Base
	basic.DataOwner
	DocumentAPI
	DocumentType         *DocumentType  `json:"document_type,omitempty" gorm:"foreignKey:DocumentTypeID"`                  // Document Type
	CitizenCountry       *Country       `json:"citizen_country,omitempty" gorm:"foreignKey:CitizenCountryID"`              // Citizen Country
	BirthCountry         *Country       `json:"birth_country,omitempty" gorm:"foreignKey:BirthCountryID"`                  // Birth Country Id
	HolderNamePrefix     *NamePrefix    `json:"holder_name_prefix,omitempty" gorm:"foreignKey:HolderNamePrefixID"`         // Holder Name Prefix
	DocumentIssueAddress *Address       `json:"document_issue_address,omitempty" gorm:"foreignKey:DocumentIssueAddressID"` // Document Issue Address
	DocumentAsset        *DocumentAsset `json:"document_asset" gorm:"foreignKey:DocumentID" swaggerignore:"true"`
}

// DocumentAPI for secure request body API
type DocumentAPI struct {
	DocumentIssueAuthority *string      `json:"document_issue_authority,omitempty" gorm:"type:varchar(64)"`                               // Document Issue Authority
	DocumentIssueLocation  *string      `json:"document_issue_location,omitempty" gorm:"type:varchar(64)"`                                // Document Issue Location
	DocumentIssueAddressID *uuid.UUID   `json:"document_issue_address_id,omitempty" gorm:"type:varchar(36)"`                              // Document Issue Address Id
	DocumentNumber         *string      `json:"document_number,omitempty" gorm:"type:varchar(36);not null"`                               // Document Number
	DocumentTypeID         *uuid.UUID   `json:"document_type_id,omitempty" gorm:"type:varchar(36);not null"`                              // Document Type Id
	EffectiveDate          *strfmt.Date `json:"effective_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"` // Effective Date
	ExpireDate             *strfmt.Date `json:"expire_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`    // Expire Date
	HolderNamePrefixID     *uuid.UUID   `json:"holder_name_prefix_id,omitempty" gorm:"type:varchar(36)"`                                  // Holder Name Prefix Id
	HolderGivenName        *string      `json:"holder_given_name,omitempty" gorm:"type:varchar(128)"`                                     // Holder Given Name
	HolderMiddleName       *string      `json:"holder_middle_name,omitempty" gorm:"type:varchar(128)"`                                    // Holder Middle Name
	HolderSurname          *string      `json:"holder_surname,omitempty" gorm:"type:varchar(128)"`                                        // Holder Surname
	BirthDate              *strfmt.Date `json:"birth_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`     // Birth Date
	BirthCountryID         *uuid.UUID   `json:"birth_country_id,omitempty" gorm:"type:varchar(36)"`                                       // Birth Country Id
	BirthPlace             *string      `json:"birth_place,omitempty" gorm:"type:varchar(64)"`                                            // Birth Place
	CitizenCountryID       *uuid.UUID   `json:"citizen_country_id,omitempty" gorm:"type:varchar(36)"`                                     // Citizen Country Id
	ContactName            *string      `json:"contact_name,omitempty" gorm:"type:varchar(256)"`                                          // Contact Name
}
