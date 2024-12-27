package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// ConvenienceStore model
type ConvenienceStore struct {
	basic.Base
	basic.DataOwner
	ConvenienceStoreAPI
	ConvenienceStoreTranslation *ConvenienceStoreTranslation `json:"convenience_store_translation,omitempty"`
}

// ConvenienceStoreAPI model
type ConvenienceStoreAPI struct {
	ConvenienceStoreCode *string `json:"convenience_store_code,omitempty" gorm:"type:varchar(36);uniqueIndex:idx_convenience_store_code_deleted_at,where:deleted_at is null;not null"`
	ConvenienceStoreName *string `json:"convenience_store_name,omitempty" gorm:"type:varchar(256);not null;index:unique_convenience_store__convenience_store_name,unique,where:deleted_at is null"`
}
