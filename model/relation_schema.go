package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

type RelationSchema struct {
	basic.Base
	ColumnSource *string `json:"column_source,omitempty" gorm:"index:idx_relation_schema__column_unique,unique;not null;" swaggertype:"string"`
	TableSource  *string `json:"table_source,omitempty" gorm:"index:idx_relation_schema__column_unique,unique;not null;" swaggertype:"string"`
	UsedByColumn *string `json:"used_by_column,omitempty" gorm:"index:idx_relation_schema__column_unique,unique;not null;" swaggertype:"string"`
	UsedByTable  *string `json:"used_by_table,omitempty" gorm:"index:idx_relation_schema__column_unique,unique;not null;" swaggertype:"string"`
}
