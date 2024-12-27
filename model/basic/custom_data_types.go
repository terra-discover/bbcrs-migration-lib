package basic

import (
	"database/sql/driver"
	"errors"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Custom data types
type CustomDataTypes []map[string]interface{}

func (p *CustomDataTypes) Value() (driver.Value, error) {
	return lib.JSONMarshal(p)
}

func (p *CustomDataTypes) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return lib.JSONUnmarshal(b, &p)
}

func (CustomDataTypes) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// returns different database type based on driver name
	switch db.Dialector.Name() {
	case "mysql", "sqlite":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

func (CustomDataTypes) GormDataType() string {
	return "json"
}
