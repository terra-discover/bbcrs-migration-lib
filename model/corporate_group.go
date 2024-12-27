package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateGroup Corporate Group
type CorporateGroup struct {
	basic.Base
	basic.DataOwner
	CorporateGroupAPI
	Office *Office `json:"office,omitempty" gorm:"foreignKey:OfficeID;references:ID"` // Office

}

// CorporateGroupAPI Corporate Group API
type CorporateGroupAPI struct {
	CorporateGroupCode *string    `json:"corporate_group_code,omitempty" gorm:"type:varchar(256);not null"` // The code of the corporate group | Rule = Must be unique
	CorporateGroupName *string    `json:"corporate_group_name,omitempty" gorm:"type:varchar(512);not null"` // The name of the corporate group | Rule = Must be unique
	OfficeID           *uuid.UUID `json:"office_id,omitempty" swaggertype:"string" format:"uuid"`           // The reference to the office. This indicates the grouping is based on an office
	IsSelfBookingTool  *bool      `json:"is_self_booking_tool,omitempty"`                                   // Indicates whether this corporate group is for corporate self booking tool | Rule =  false: This corporate group is not for corporate self booking tool (this is default). true: This corporate group is for corporate self booking tool.
	Description        *string    `json:"description,omitempty" gorm:"type:text"`                           // The description of the corporate group.
}
