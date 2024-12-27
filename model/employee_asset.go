package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EmployeeAsset struct
type EmployeeAsset struct {
	basic.Base
	basic.DataOwner
	EmployeeAssetAPI
	Employee              *Employee              `json:"employee,omitempty" swaggerignore:"true"`                                                                                                                          // employee
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                                                                                 // multimedia description
	EmployeeID            *uuid.UUID             `json:"employee_id,omitempty" gorm:"type:varchar(36);index:employee_asset_multimedia_unique,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // employee id
}

// EmployeeAssetAPI struct writable by request body
type EmployeeAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);index:employee_asset_multimedia_unique,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
