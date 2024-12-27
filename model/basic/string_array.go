package basic

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// StringArray data type
type StringArray []string

func (t *StringArray) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		oti, ok := value.([]uint8)
		if !ok {
			return errors.New(fmt.Sprint("wrong type : ", value))
		}
		val = string(oti)
	}
	*t = StringArray(strings.Split(string(val), "|"))
	return nil
}

func (t StringArray) Value() (driver.Value, error) {
	//this check is here if you don't want to save an empty string
	if len(t) == 0 {
		return nil, nil
	}
	return []byte(strings.Join(t, "|")), nil
}

func (StringArray) GormDataType() string {
	return "text"
}

func (StringArray) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// returns different database type based on driver name
	switch db.Dialector.Name() {
	case "postgres", "sqlite":
		return "text"
	}
	return ""
}
