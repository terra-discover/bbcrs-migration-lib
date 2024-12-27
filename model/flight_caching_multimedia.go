package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type FlightCachingMultimedia struct {
	basic.Base
	basic.DataOwner
	FlightCachingMultimediaAPI
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`
}

type FlightCachingMultimediaAPI struct {
	SessionID               *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);not null"`
	MultimediaType          *string    `json:"multimedia_type,omitempty" gorm:"type:varchar(100);not null"`
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null"`
}

type FlightCachingMultimediaType string

const (
	MultimediaTypeEticket    FlightCachingMultimediaType = "eticket"
	MultimediaTypeEitinerary FlightCachingMultimediaType = "eitinerary"
)

func (fcm FlightCachingMultimediaType) String() string {
	return string(fcm)
}
