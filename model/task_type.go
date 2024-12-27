package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// TaskType Task Type
type TaskType struct {
	basic.Base
	basic.DataOwner
	TaskTypeAPI
	TaskTypeTranslation *TaskTypeTranslation `json:"task_type_translation,omitempty"` // Task Type Translation
}

// TaskTypeAPI Task Type API
type TaskTypeAPI struct {
	TaskTypeCode *int    `json:"task_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`           // Task Type Code
	TaskTypeName *string `json:"task_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"General"` // Task Type Name
}

// Seed data
func (s TaskType) Seed() *[]TaskType {
	data := []TaskType{}
	items := []string{
		"General",
		"Prepare Proposal",
		"Book Flight",
		"Book Hotel",
		"Issue Flight Ticket",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, TaskType{
			TaskTypeAPI: TaskTypeAPI{
				TaskTypeCode: &code,
				TaskTypeName: &name,
			},
		})
	}

	return &data
}
