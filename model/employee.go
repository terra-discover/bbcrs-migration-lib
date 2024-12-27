package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// Employee Employee
type Employee struct {
	basic.Base
	basic.DataOwner
	EmployeeAPI
	UserAccount              *UserAccount              `json:"user_account,omitempty" gorm:"foreignKey:UserAccountID;references:ID"`
	JobTitle                 *JobTitle                 `json:"job_title,omitempty" gorm:"foreignKey:JobTitleID;references:ID"`
	Division                 *Division                 `json:"division,omitempty" gorm:"foreignKey:DivisionID;references:ID"`
	Office                   *Office                   `json:"office,omitempty" gorm:"foreignKey:OfficeID;references:ID"`
	Person                   *Person                   `json:"person,omitempty" gorm:"foreignKey:PersonID;references:ID"`
	CorporateEmployee        *CorporateEmployee        `json:"corporate_employee,omitempty" gorm:"foreignKey:ID;references:EmployeeID" swaggerignore:"true"`
	CorporateUser            *CorporateUser            `json:"corporate_user,omitempty" gorm:"foreignKey:UserAccountID;references:UserAccountID" swaggerignore:"true"`
	EmployeeAsset            *EmployeeAsset            `json:"employee_asset,omitempty" swaggerignore:"true"`
	AgentEmployee            *AgentEmployee            `json:"agent_employee,omitempty" gorm:"foreignKey:EmployeeID;references:ID" swaggerignore:"true"`
	AgentUser                *AgentUser                `json:"agent_user,omitempty" gorm:"foreignKey:UserAccountID" swaggerignore:"true"`
	AgentTeamConsultant      *AgentTeamConsultant      `json:"agent_team_consultant,omitempty" swaggerignore:"true"`
	AgentCorporateConsultant *AgentCorporateConsultant `json:"agent_corporate_consultant,omitempty" swaggerignore:"true"`
	EmployeeAdditionalRole   *EmployeeAdditionalRole   `json:"employee_additional_role,omitempty" gorm:"foreignKey:EmployeeID;references:ID" swaggerignore:"true" `
}

// EmployeeAPI EmployeeAPI
type EmployeeAPI struct {
	EmployeeNumber *string           `json:"employee_number,omitempty" gorm:"type:varchar(36);not null;"`
	PersonID       *uuid.UUID        `json:"person_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggerignore:"true"`
	UserAccountID  *uuid.UUID        `json:"user_account_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	JobTitleID     *uuid.UUID        `json:"job_title_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	HireDate       *strfmt.DateTime  `json:"hire_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	DivisionID     *uuid.UUID        `json:"division_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	OfficeID       *uuid.UUID        `json:"office_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	EmployeeAsset  *EmployeeAssetAPI `json:"employee_asset,omitempty" gorm:"-"`
}

// Seed Employee
func (e *Employee) Seed(userID uuid.UUID) *[]Employee {
	data := []Employee{}

	if lib.IsEmptyUUID(userID) {
		userID = uuid.New()
	}

	data = append(data, Employee{
		EmployeeAPI: EmployeeAPI{
			EmployeeNumber: lib.Strptr("1"),
			UserAccountID:  &userID,
			PersonID:       &userID,
		},
	})

	return &data
}
