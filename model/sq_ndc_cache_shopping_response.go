package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// SqNdcCacheShoppingResponse table struct
type SqNdcCacheShoppingResponse struct {
	basic.Base
	basic.DataOwner
	SessionID          *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`
	ShoppingResponseID *string    `json:"shopping_response_id,omitempty" gorm:"type:varchar(60)" default:"SP2F-1451030339040814687"`
	ShoppingType       *string    `json:"shopping_type,omitempty" gorm:"type:varchar(15);default:'flight'" default:"flight" example:"flight" `
	IsOffer            *bool      `json:"is_offer,omitempty"`
}

// GetCurrentShoppingResponseID by query filter
func (data *SqNdcCacheShoppingResponse) GetCurrentShoppingResponseID(tx *gorm.DB, sessionID *uuid.UUID, shoppingType *string) {
	tx.Where(tx.Where(SqNdcCacheShoppingResponse{
		SessionID:    sessionID,
		ShoppingType: shoppingType,
	})).Order(`updated_at desc`).Take(&data)
}
