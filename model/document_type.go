package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DocumentType Document Type
type DocumentType struct {
	basic.Base
	basic.DataOwner
	DocumentTypeAPI
	DocumentTypeTranslation *DocumentTypeTranslation `json:"document_type_translation,omitempty"`
}

// DocumentTypeAPI Document Type API
type DocumentTypeAPI struct {
	DocumentTypeCode *int    `json:"document_type_code,omitempty" gorm:"type:smallint;not null;index:idx_document_type_code_deleted_at,unique,where:deleted_at is null" example:"1"`        // Document Type Code
	DocumentTypeName *string `json:"document_type_name,omitempty" gorm:"type:varchar(256);not null;index:idx_document_type_name_deleted_at,unique,where:deleted_at is null" example:"Visa"` // Document Type Name
}

// // Seed Document Type
// func (documentType *DocumentType) Seed() *DocumentType {
// 	documentType.DocumentTypeCode = lib.Intptr(2)
// 	documentType.DocumentTypeName = lib.Strptr("Passport")
// 	return documentType
// }

// Seed for seeder db
func (documentType *DocumentType) Seed() *[]DocumentType {
	name := []string{"Visa", "Passport", "Military identification", "Drivers license", "National identity document", "Vaccination certificate",
		"Alien registration number", "Insurance policy number", "Tax exemption number", "Vehicle registration/license number", "Border crossing card",
		"Refugee travel document", "Pilot's license", "Permanent resident card", "Redress number", "Known traveler number", "Non-standard", "Merchant mariner",
		"Air Nexus card", "Crew member certificate", "Passport card", "Naturalization certificate", "Tax identification number"}
	j := 1
	result := []DocumentType{}
	for i := 0; i < len(name); i++ {
		if name[i] == "Tax identification number" {
			j = 1001
		}
		result = append(result, DocumentType{DocumentTypeAPI: DocumentTypeAPI{
			DocumentTypeCode: lib.Intptr(j),
			DocumentTypeName: lib.Strptr(name[i]),
		}})
		j++
	}
	return &result
}
