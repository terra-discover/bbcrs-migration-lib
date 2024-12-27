package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// GuestRoomInfoType Guest Room Info Type
type GuestRoomInfoType struct {
	basic.Base
	basic.DataOwner
	GuestRoomInfoTypeAPI
	GuestRoomInfoTypeTranslation *GuestRoomInfoTypeTranslation `json:"guest_room_info_type_translation,omitempty"` // Guest Room Info Type Translation
}

// GuestRoomInfoTypeAPI Guest Room Info Type API
type GuestRoomInfoTypeAPI struct {
	GuestRoomInfoTypeCode *int    `json:"guest_room_info_type_code,omitempty" gorm:"type:int;index:,unique,where:deleted_at is null;not null" example:"1"`                         // Guest Room Info Type Code
	GuestRoomInfoTypeName *string `json:"guest_room_info_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Accessible rooms"` // Guest Room Info Type Name
	Description           *string `json:"description,omitempty" gorm:"type:text" example:"Accessible rooms"`                                                                       // Description
}

// Seed data
func (s GuestRoomInfoType) Seed() *[]GuestRoomInfoType {
	data := []GuestRoomInfoType{}
	items := []string{
		"Accessible rooms",
		"Nonsmoking rooms",
		"Suites",
		"Bungalows and villas",
		"Floors",
		"Executive floor",
		"Rooms that work",
		"Available rooms",
		"Available suites",
		"Double bedrooms",
		"King bedrooms",
		"Total rooms",
		"Apartments",
		"Queen bedrooms",
		"Penthouses",
		"Studios",
		"First floor rooms",
		"Smoking rooms",
		"Twin bedrooms",
		"Drive up rooms",
		"Rooms with internet access",
		"Freestanding units",
		"Air conditioned guest rooms",
		"Concierge levels",
		"Condos",
		"Club levels",
		"Total available rooms and suites",
		"Total rooms and suites",
		"Employees on property",
		"Employees working for property",
		"Separate floors for women",
		"Buildings",
		"Accommodations with balcony",
		"Adjoining rooms or suites",
		"Connecting rooms or suites",
		"Family/oversized accommodations",
		"Single-bedded accommodations",
		"Cabin",
		"Cottage",
		"Loft",
		"Parlour",
		"Room",
		"Lanai",
		"Bungalow",
		"Villa",
		"Efficiency",
		"All rooms non-smoking",
		"Double double bedrooms",
		"King king bedrooms",
		"Queen queen bedrooms",
		"Twin twin bedrooms",
		"Apartment for 1",
		"Apartment for 2",
		"Apartment for 3",
		"Apartment for 4",
		"Apartment for 6",
		"1 room cabin",
		"1 bedroom cabin",
		"2 bedroom cabin",
		"Junior suite",
		"Jacuzzi suite",
		"Run of the house",
		"Large suite",
		"1 bedroom suite",
		"2 bedroom suite",
		"3 bedroom suite",
		"Villa for 1",
		"Villa for 2",
		"Villa for 3",
		"Villa for 6",
		"Villa for 8",
		"Single with pullout",
		"Business plan",
		"Business class",
		"Classic",
		"Comfort",
		"Deluxe",
		"Deluxe suite",
		"Economy",
		"Luxury",
		"Premier",
		"Standard",
		"Superior",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]

		data = append(data, GuestRoomInfoType{
			GuestRoomInfoTypeAPI: GuestRoomInfoTypeAPI{
				GuestRoomInfoTypeCode: &code,
				GuestRoomInfoTypeName: &name,
			},
		})
	}

	return &data
}
