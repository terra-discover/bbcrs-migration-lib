package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentAsset struct
type AgentAsset struct {
	basic.Base
	basic.DataOwner
	AgentAssetAPI
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`
	AgentID               *uuid.UUID             `json:"agent_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Agent id
}

// AgentAssetAPI struct
type AgentAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
