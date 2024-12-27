package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BusinessEntityDocument struct
type BusinessEntityDocument struct {
	basic.Base
	basic.DataOwner
	BusinessEntityDocumentAPI
	BusinessEntity *BusinessEntity `json:"business_entity,omitempty"`
	Document       *Document       `json:"document,omitempty"`
}

// BusinessEntityDocumentAPI for secure request body API
type BusinessEntityDocumentAPI struct {
	BusinessEntityID *uuid.UUID `json:"business_entity_id,omitempty" gorm:"type:varchar(36);not null;index:idx_business_entity_document__business_entity_id,unique,where:deleted_at is null;"`
	DocumentID       *uuid.UUID `json:"document_id,omitempty" gorm:"type:varchar(36);not null;index:idx_business_entity_document__business_entity_id,unique,where:deleted_at is null;"`
}
