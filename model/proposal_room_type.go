package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProposalRoomType Proposal Room Type
type ProposalRoomType struct {
	basic.Base
	basic.DataOwner
	ProposalRoomTypeAPI
}

// ProposalRoomTypeAPI Proposal Room Type Api
type ProposalRoomTypeAPI struct {
	ProposalID   *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`  // Proposal Id
	RoomTypeID   *uuid.UUID `json:"room_type_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"` // Room Type Id
	RoomTypeCode *string    `json:"room_type_code,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`                                                    // Room Type Code
	RoomTypeName *string    `json:"room_type_name,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`                                                    // Room Type Name
	IsAccrual    *bool      `json:"is_accrual,omitempty" gorm:"type:bool;default:false" swaggertype:"boolean"`                                                // Is Accrual
	IsRedeemable *bool      `json:"is_redeemable,omitempty" gorm:"type:bool;default:false" swaggertype:"boolean"`                                             // Is Redeemable
}
