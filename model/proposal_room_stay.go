package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProposalRoomStay model
type ProposalRoomStay struct {
	basic.Base
	basic.DataOwner
	ProposalRoomStayAPI
	Proposal *Proposal `json:"proposal"`
	RoomStay *RoomStay `json:"room_stay"`
}

// ProposalRoomStayAPI struct
type ProposalRoomStayAPI struct {
	ProposalID *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);uniqueIndex:,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`  // Porposal ID
	RoomStayID *uuid.UUID `json:"room_stay_id,omitempty" gorm:"type:varchar(36);uniqueIndex:,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Room Stay ID
}
