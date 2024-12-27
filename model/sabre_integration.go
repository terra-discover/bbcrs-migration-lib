package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type SabreIntegration struct {
	basic.Base
	basic.DataOwner
	SessionID      *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`
	Username       *string    `json:"username,omitempty" gorm:"type:varchar(256)"`
	Password       *string    `json:"password,omitempty" gorm:"type:varchar(256)"`
	Organization   *string    `json:"organization,omitempty" gorm:"type:varchar(256)"`
	Domain         *string    `json:"domain,omitempty" gorm:"type:varchar(256)"`
	PCC            *string    `json:"pcc,omitempty" gorm:"type:varchar(256)"`
	IPCC           *string    `json:"ipcc,omitempty" gorm:"type:varchar(256)"`
	PrinterAddress *string    `json:"printer_address,omitempty" gorm:"type:varchar(256)"`
}

// func (b *SabreIntegration) Seed() *SabreIntegration {
// 	data := SabreIntegration{
// 		Username:       lib.Strptr("666"),
// 		Password:       lib.Strptr("wisata40"),
// 		Organization:   lib.Strptr("D8UD"),
// 		Domain:         lib.Strptr("AA"),
// 		PCC:            lib.Strptr("F9Q8"),
// 		IPCC:           lib.Strptr("D8UD"),
// 		PrinterAddress: lib.Strptr("A9EFD9"),
// 	}
// 	_ = lib.Merge(data, &b)
// 	return b

// }
