package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FormType Form Type
type FormType struct {
	basic.Base
	basic.DataOwner
	FormTypeAPI
	FormTypeTranslation *FormTypeTranslation `json:"form_type_translation,omitempty"`
}

// FormTypeAPI Form Type API
type FormTypeAPI struct {
	FormTypeCode *int    `json:"form_type_code,omitempty" gorm:"type:smallint;not null;index:idx_form_type_code_deleted_at,unique,where:deleted_at is null" example:"1"`          // Form Type Code
	FormTypeName *string `json:"form_type_name,omitempty" gorm:"type:varchar(256);not null;index:idx_form_type_name_deleted_at,unique,where:deleted_at is null" example:"Survey"` // Form Type Name
}

// Seed FormType
func (formType *FormType) Seed() *FormType {
	seed := FormType{
		FormTypeAPI: FormTypeAPI{
			FormTypeCode: lib.Intptr(1),
			FormTypeName: lib.Strptr("Survey"),
		},
	}
	_ = lib.Merge(seed, &formType)
	return formType
}

// FormTypes struct
type FormTypes []FormType

// Seed for seeder
func (formTypes *FormTypes) Seed() *FormTypes {
	name := []string{"General Inquiry", "Feedback", "Survey"}
	for i := 0; i < len(name); i++ {
		*formTypes = append(*formTypes, FormType{FormTypeAPI: FormTypeAPI{
			FormTypeCode: lib.Intptr(i + 1),
			FormTypeName: lib.Strptr(name[i]),
		}})
	}
	return formTypes
}
