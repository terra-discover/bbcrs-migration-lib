package model

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BusinessEntity struct
type BusinessEntity struct {
	basic.Base
	basic.DataOwner
	BusinessEntityAPI
	Language              *Language               `json:"language,omitempty"`  // Language
	Currency              *Currency               `json:"currency,omitempty"`  // Currency
	Company               *Company                `json:"company,omitempty"`   // Company
	Industry              *Industry               `json:"industry,omitempty"`  // Industry
	TimeZone              *TimeZone               `json:"time_zone,omitempty"` //Timezone
	BusinessEntityAddress []BusinessEntityAddress `json:"business_entity_address,omitempty" gorm:"foreignKey:BusinessEntityID;references:ID"`
}

// BusinessEntityAPI for secure request body API
type BusinessEntityAPI struct {
	EffectiveDate           *strfmt.DateTime `json:"effective_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"` // Effective Date
	ExpireDate              *strfmt.DateTime `json:"expire_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`    // Expire Date
	RegisteredByID          *uuid.UUID       `json:"registered_by_id,omitempty" gorm:"type:varchar(36)"`                                       // Registered By user Id
	RegisteredAt            *strfmt.DateTime `json:"registered_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`  // Registered At
	RegistrationReason      *string          `json:"registration_reason,omitempty" gorm:"type:text"`                                           // Registration Reason
	ApprovedByID            *uuid.UUID       `json:"approved_by_id,omitempty" gorm:"type:varchar(36)"`                                         // Approved By user Id
	ApprovedAt              *strfmt.DateTime `json:"approved_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`    // Approved At
	ApprovalReason          *string          `json:"approval_reason,omitempty" gorm:"type:text"`                                               // Approval Reason
	RejectedByID            *uuid.UUID       `json:"rejected_by_id,omitempty" gorm:"type:varchar(36)"`                                         // Rejected By user Id
	RejectedAt              *strfmt.DateTime `json:"rejected_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`    // Rejected At
	RejectionReason         *string          `json:"rejection_reason,omitempty" gorm:"type:text"`                                              // Rejected Reason
	TerminatedByID          *uuid.UUID       `json:"terminated_by_id,omitempty" gorm:"type:varchar(36)"`                                       // Terminated By user Id
	TerminatedAt            *strfmt.DateTime `json:"terminated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`  // Terminated At
	TerminationReason       *string          `json:"termination_reason,omitempty" gorm:"type:text"`                                            // Terminated Reason
	LanguageID              *uuid.UUID       `json:"language_id,omitempty" gorm:"type:varchar(36)"`                                            // Language Id
	CurrencyID              *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36)"`                                            // Currency Id
	TaxCategoryID           *uuid.UUID       `json:"tax_category_id,omitempty" gorm:"type:varchar(36)"`                                        // Tax Category Id
	DistanceUnitOfMeasureID *uuid.UUID       `json:"distance_unit_of_measure_id,omitempty" gorm:"type:varchar(36)"`                            // Distance Unit Of Measure Id
	AreaUnitOfMeasureID     *uuid.UUID       `json:"area_unit_of_measure_id,omitempty" gorm:"type:varchar(36)"`                                // Area Unit Of Measure Id
	WeightUnitOfMeasureID   *uuid.UUID       `json:"weight_unit_of_measure_id,omitempty" gorm:"type:varchar(36)"`                              // Weight Unit Of Measure Id
	CompanyID               *uuid.UUID       `json:"company_id,omitempty" gorm:"type:varchar(36)"`                                             // Company Id
	CompanyOther            *string          `json:"company_other,omitempty" gorm:"type:varchar(256)"`                                         // Company Other
	IndustryID              *uuid.UUID       `json:"industry_id,omitempty" gorm:"type:varchar(36)"`                                            // Industry Id
	TimeZoneID              *uuid.UUID       `json:"time_zone_id,omitempty" gorm:"type:varchar(36)"`                                           // Time Zone ID
}

// Seed BusinessEntity
func (b *BusinessEntity) Seed(agentID uuid.UUID) *[]BusinessEntity {
	now := strfmt.DateTime(time.Now().UTC())

	data := []BusinessEntity{}

	if lib.IsEmptyUUID(agentID) {
		agentID = uuid.New()
	}

	businessEntityID := agentID
	initial := BusinessEntity{}
	initial.ID = &businessEntityID
	initial.TimeZoneID = &businessEntityID
	initial.CreatedAt = &now
	data = append(data, initial)

	return &data
}
