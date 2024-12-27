package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Language Language
type Language struct {
	basic.Base
	basic.DataOwner
	LanguageAPI
	LanguageTranslation *LanguageTranslation `json:"language_translation,omitempty"` // language translation
	LanguageAsset       *LanguageAsset       `json:"language_asset,omitempty"`       // language asset
}

// LanguageAPI Language API
type LanguageAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);index:,unique,where:deleted_at is null;not null" example:"en"`                                                                             // 2 Characters language code
	LanguageName       *string `json:"language_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"English"`                                                                      // International Language Name
	LanguageNativeName *string `json:"language_native_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"English"`                                                               // Language Foreign Name
	LanguageAlpha3Code *string `json:"language_alpha_3_code,omitempty" gorm:"column:language_alpha_3_code;type:varchar(3);index:,unique,where:deleted_at is null and length(language_alpha_3_code) > 0;not null" example:"eng"` // 3 characters language codes
	// TODO: must review in master-service
	// LanguageAsset      *LanguageAssetAPI `json:"language_asset,omitempty" gorm:"-"`                                                                                                                                                       // language asset
}

// Seed Language
func (l *Language) Seed() *[]Language {
	data := []Language{}

	initial := Language{}
	initial.LanguageCode = lib.Strptr("id")
	initial.LanguageName = lib.Strptr("Indonesia")
	initial.LanguageAlpha3Code = lib.Strptr("idn")
	initial.LanguageNativeName = lib.Strptr("Bahasa Indonesia")
	data = append(data, initial)

	initial2 := Language{}
	initial2.LanguageCode = lib.Strptr("en")
	initial2.LanguageName = lib.Strptr("English")
	initial2.LanguageAlpha3Code = lib.Strptr("eng")
	initial2.LanguageNativeName = lib.Strptr("English")
	data = append(data, initial2)

	return &data
}
