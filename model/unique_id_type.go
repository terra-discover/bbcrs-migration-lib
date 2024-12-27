package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// UniqueIDType Unique ID Type
type UniqueIDType struct {
	basic.Base
	basic.DataOwner
	UniqueIDTypeAPI
	UniqueIDTypeTranslation *UniqueIDTypeTranslation `json:"unique_id_type_translation,omitempty"` // Unique ID Type Translation
}

// UniqueIDTypeAPI Unique ID Type API
type UniqueIDTypeAPI struct {
	UniqueIDTypeCode *int    `json:"unique_id_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`            // Unique ID Type Code
	UniqueIDTypeName *string `json:"unique_id_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Customer"` // Unique ID Type Name
}

// Seed data
func (r UniqueIDType) Seed() *[]UniqueIDType {
	data := []UniqueIDType{}
	items := []string{
		"Customer",
		"CRO (Customer Reservations Office)",
		"Corporation representative",
		"Company",
		"Travel agency",
		"Airline",
		"Wholesaler",
		"Car rental",
		"Group",
		"Hotel",
		"Tour operator",
		"Cruise line",
		"Internet broker",
		"Reservation",
		"Cancellation",
		"Reference",
		"Meeting planning agency",
		"Other",
		"Insurance agency",
		"Insurance agent",
		"Profile",
		"ERSP (Electronic reservation service provider)",
		"Provisional reservation",
		"Travel Agent PNR",
		"Associated reservation",
		"Associated itinerary reservation",
		"Associated shared reservation",
		"Alliance",
		"Booking agent",
		"Ticket",
		"Divided reservation",
		"Merchant",
		"Acquirer",
		"Master reference",
		"Purged master reference",
		"Parent reference",
		"Child reference",
		"Linked reference",
		"Contract",
		"Confirmation number",
		"Fare quote",
		"Reissue/refund quote",
		"Ground transportation supplier",
		"EMD",
		"TMS",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, UniqueIDType{
			UniqueIDTypeAPI: UniqueIDTypeAPI{
				UniqueIDTypeCode: &code,
				UniqueIDTypeName: &name,
			},
		})
	}

	return &data
}
