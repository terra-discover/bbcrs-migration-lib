package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReferralSource Referral Source
type ReferralSource struct {
	basic.Base
	basic.DataOwner
	ReferralSourceAPI
	ReferralSourceTranslation *ReferralSourceTranslation `json:"referral_source_translation,omitempty"`
}

// ReferralSourceAPI Referral Source API
type ReferralSourceAPI struct {
	ReferralSourceName *string `json:"referral_source_name,omitempty" gorm:"type:varchar(64);not null;index:,unique,where:deleted_at is null" example:"Referral Source Name"` // Referral Source Name
}

// Seed ReferralSource
func (referralSource *ReferralSource) Seed() *ReferralSource {
	seed := ReferralSource{
		ReferralSourceAPI: ReferralSourceAPI{
			ReferralSourceName: lib.Strptr("Referral Source Name"),
		},
	}
	_ = lib.Merge(seed, &referralSource)
	return referralSource
}
