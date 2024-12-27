package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyAccount struct
type LoyaltyAccount struct {
	basic.Base
	basic.DataOwner
	LoyaltyAccountAPI
	EnrollmentMethod *EnrollmentMethod `json:"enrollment_method,omitempty"` // enrollment_method
	LoyaltyProgram   *LoyaltyProgram   `json:"loyalty_program,omitempty"`   // Loyalty Program
}

// LoyaltyAccountAPI for secure request body API
type LoyaltyAccountAPI struct {
	MemberNumber             *string      `json:"member_number,omitempty" gorm:"type:varchar(36);not null"`
	LoyaltyProgramID         *uuid.UUID   `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36)"`
	ProfileStatus            *int         `json:"profile_status,omitempty" gorm:"type:smallint"`
	UserAccountID            *uuid.UUID   `json:"user_account_id,omitempty" gorm:"type:varchar(36)"`
	LoyaltyLevelID           *uuid.UUID   `json:"loyalty_level_id,omitempty" gorm:"type:varchar(36)"`
	LoyaltyLevelExpiryDate   *strfmt.Date `json:"loyalty_level_expiry_date,omitempty" gorm:"type:timestamptz"`
	LoyaltyLevelRewardAmount *float64     `json:"loyalty_level_reward_amount,omitempty"`
	EnrollmentMethodID       *uuid.UUID   `json:"enrollment_method_id,omitempty" gorm:"type:varchar(36)"`
	SignUpDate               *strfmt.Date `json:"sign_up_date,omitempty" gorm:"type:timestamptz"`
	EffectiveDate            *strfmt.Date `json:"effective_date,omitempty" gorm:"type:timestamptz"`
	ExpireDate               *strfmt.Date `json:"expire_date,omitempty" gorm:"type:timestamptz"`
	ApprovedByID             *uuid.UUID   `json:"approved_by_id,omitempty" gorm:"type:varchar(36)"`
	ApprovedAt               *strfmt.Date `json:"approved_at,omitempty" gorm:"type:timestamptz"`
	ApprovalReason           *string      `json:"approval_reason,omitempty"`
	RejectedByID             *uuid.UUID   `json:"rejected_by_id,omitempty" gorm:"type:varchar(36)"`
	RejectedAt               *strfmt.Date `json:"rejected_at,omitempty" gorm:"type:timestamptz"`
	RejectionReason          *string      `json:"rejection_reason,omitempty"`
	TerminatedByID           *uuid.UUID   `json:"terminated_by_id,omitempty" gorm:"type:varchar(36)"`
	TerminatedAt             *strfmt.Date `json:"terminated_at,omitempty" gorm:"type:timestamptz"`
	TerminationReason        *string      `json:"termination_reason,omitempty"`
}

// LoyaltyAccountData model
type LoyaltyAccountData struct {
	LoyaltyProgramName *string `json:"loyalty_program_name,omitempty" gorm:"type:varchar(255)"`
	LoyaltyProgramCode *string `json:"loyalty_program_code,omitempty"`
	LoyaltyAccountID   *string `json:"loyalty_account_id,omitempty" gorm:"type:varchar(255)"`
	ProductTypeName    *string `json:"type_of_travel_program,omitempty" gorm:"type:varchar(256)"`
	LoyaltyAccountDataAPI
}

// LoyaltyAccountDataAPI model
type LoyaltyAccountDataAPI struct {
	ProductTypeID    *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36)"`
	LoyaltyProgramID *uuid.UUID `json:"loyalty_program_id,omitempty" gorm:"type:varchar(36)" swaggerignore:"true"`
	LoyaltyNumber    *string    `json:"loyalty_number,omitempty" gorm:"type:varchar(255)"`
}
