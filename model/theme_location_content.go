package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ThemeLocationContent struct
type ThemeLocationContent struct {
	basic.Base
	basic.DataOwner
	ThemeLocationContentAPI
}

// ThemeLocationContentAPI for secure request body API
type ThemeLocationContentAPI struct {
	ThemeLocationID      *uuid.UUID `json:"theme_location_id,omitempty" swaggertype:"string" format:"uuid"`
	ContentDescriptionID *uuid.UUID `json:"content_description_id,omitempty" swaggertype:"string" format:"uuid"`
}
