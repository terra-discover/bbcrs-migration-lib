package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// BonusProgram model
type BonusProgram struct {
	basic.Base
	basic.DataOwner
	BonusProgramAPI
	BonusProgramTranslation *BonusProgramTranslation `json:"bonus_program_translation,omitempty"`
}

// BonusProgramAPI model
type BonusProgramAPI struct {
	BonusProgramCode *string `json:"bonus_program_code,omitempty" gorm:"type:varchar(36);uniqueIndex:idx_bonus_program_code_deleted_at,where:deleted_at is null;not null"`
	BonusProgramName *string `json:"bonus_program_name,omitempty" gorm:"type:varchar(256);not null;index:unique_bonus_program__bonus_program_name,unique,where:deleted_at is null"`
	Description      *string `json:"description,omitempty" gorm:"type:text"`
}
