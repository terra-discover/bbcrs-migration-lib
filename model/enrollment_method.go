package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EnrollmentMethod Enrollment Method
type EnrollmentMethod struct {
	basic.Base
	basic.DataOwner
	EnrollmentMethodAPI
	EnrollmentMethodTranslation *EnrollmentMethodTranslation `json:"enrollment_method_translation,omitempty"`
}

// EnrollmentMethodAPI Enrollment Method API
type EnrollmentMethodAPI struct {
	EnrollmentMethodCode *int    `json:"enrollment_method_code,omitempty" gorm:"type:smallint;not null;index:idx_enrollment_method_code_deleted_at,unique,where:deleted_at is null" example:"1"`             // Enrollment Method Code
	EnrollmentMethodName *string `json:"enrollment_method_name,omitempty" gorm:"type:varchar(256);not null;index:idx_enrollment_method_name_deleted_at,unique,where:deleted_at is null" example:"Admin web"` // Enrollment Method Name
}

// Seed Enrollment Method
func (enrollmentMethod *EnrollmentMethod) Seed() *EnrollmentMethod {
	seed := EnrollmentMethod{
		EnrollmentMethodAPI: EnrollmentMethodAPI{
			EnrollmentMethodCode: lib.Intptr(1),
			EnrollmentMethodName: lib.Strptr("Admin web"),
		},
	}
	_ = lib.Merge(seed, &enrollmentMethod)
	return enrollmentMethod
}
