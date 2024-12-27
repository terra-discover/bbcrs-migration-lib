package model

import (
	"strconv"
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// CorporateRatingTypeLevel Corporate Rating Type Level
type CorporateRatingTypeLevel struct {
	basic.Base
	basic.DataOwner
	CorporateRatingTypeLevelAPI
	Translation         *CorporateRatingTypeLevelTranslation `json:"translation,omitempty"`           // Corporate Rating Type Level Translation
	CorporateRatingType *CorporateRatingType                 `json:"corporate_rating_type,omitempty"` // Corporate Rating Type
}

// CorporateRatingTypeLevelAPI Corporate Rating Type Level API
type CorporateRatingTypeLevelAPI struct {
	CorporateRatingTypeLevelCode *int       `json:"corporate_rating_type_level_code,omitempty" example:"1" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null"`       // Corporate Rating Type Level Code
	CorporateRatingTypeLevelName *string    `json:"corporate_rating_type_level_name,omitempty" example:"Low" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Corporate Rating Type Level Name
	CorporateRatingTypeID        *uuid.UUID `json:"corporate_rating_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`                            // Corporate Rating Type ID
	Rating                       *float64   `json:"rating,omitempty" example:"0"`                                                                                                      // Rating
}

// Seed data
func (s CorporateRatingTypeLevel) Seed() *[]CorporateRatingTypeLevel {
	data := []CorporateRatingTypeLevel{}
	items := []string{
		"1|Low|1",
		"2|Medium|2",
		"3|High|3",
	}

	for i := range items {
		contents := strings.Split(items[i], "|")
		code, _ := strconv.Atoi(contents[0])
		rating, _ := strconv.Atoi(contents[2])
		var name string = contents[1]
		data = append(data, CorporateRatingTypeLevel{
			CorporateRatingTypeLevelAPI: CorporateRatingTypeLevelAPI{
				CorporateRatingTypeLevelCode: &code,
				CorporateRatingTypeLevelName: &name,
				CorporateRatingTypeID:        lib.UUIDPtr(uuid.New()),
				Rating:                       lib.Float64ptr(float64(rating)),
			},
		})
	}

	return &data
}
