package migration

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Test__ModelMigrations(t *testing.T) {
	// 1. check duplicate model
	checkDuplicateModel(t)
	// 2. check valid model
	checkValidModel(t)
}

func checkDuplicateModel(t *testing.T) {
	/*
		Fix error "no such table" in gorm sqlite using in-memory database.
		":memory:" will open a separate database for each connection. Use "file::memory:?cache=shared" instead.
		source:
		- https://stackoverflow.com/a/62994222
		- https://github.com/mattn/go-sqlite3/issues/223#issuecomment-573539458
	*/
	sharedMemory := "file::memory:?cache=shared"

	db, err := gorm.Open(sqlite.Open(sharedMemory), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
	})
	utils.AssertEqual(t, nil, err, "mock db")

	t.Cleanup(func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	})

	// Migrating
	err = db.AutoMigrate(ModelMigrations...)
	utils.AssertEqual(t, nil, err, "mock db")

	listUniqueTable := []string{}

	arrMessage := []string{}

loopModelMigrations:
	for idx, model := range ModelMigrations {
		assignModel := model
		mod := db.Model(assignModel).Take(assignModel)
		table := mod.Statement.Table
		if lib.IsEmptyStr(table) {
			arrMessage = append(arrMessage, fmt.Sprintf("table is empty on index: %d", idx))
			continue loopModelMigrations
		}

		isMatch := false

	loopListUniqueTable:
		for _, uqTable := range listUniqueTable {
			if uqTable == table {
				arrMessage = append(arrMessage, fmt.Sprintf("duplicate model declared of table %s", uqTable))
				break loopListUniqueTable
			}
		}

		if !isMatch {
			listUniqueTable = append(listUniqueTable, table)
		}
	}

	// Log err
	if len(arrMessage) > 0 {
		message := strings.Join(arrMessage, ";\n")
		t.Errorf("ERROR checkDuplicateModel: \n %s", message)
	}
}

func checkValidModel(t *testing.T) {
	/*
		Fix error "no such table" in gorm sqlite using in-memory database.
		":memory:" will open a separate database for each connection. Use "file::memory:?cache=shared" instead.
		source:
		- https://stackoverflow.com/a/62994222
		- https://github.com/mattn/go-sqlite3/issues/223#issuecomment-573539458
	*/
	sharedMemory := "file::memory:?cache=shared"

	db, err := gorm.Open(sqlite.Open(sharedMemory), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
	})
	utils.AssertEqual(t, nil, err, "mock db")

	t.Cleanup(func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	})

	// Migrating
	err = db.AutoMigrate(ModelMigrations...)
	utils.AssertEqual(t, nil, err, "mock db")

	arrMessage := []string{}

loopModelMigrations:
	for idx, model := range ModelMigrations {
		assignModel := model
		mod := db.Model(assignModel).Take(assignModel)
		if mod.Error != nil && mod.Error != gorm.ErrRecordNotFound {
			arrMessage = append(arrMessage, fmt.Sprintf("error getting table index: %d , message: %s", idx, mod.Error.Error()))
			continue loopModelMigrations
		}

		table := mod.Statement.Table
		if lib.IsEmptyStr(table) {
			arrMessage = append(arrMessage, fmt.Sprintf("table is empty on index: %d", idx))
			continue loopModelMigrations
		}
	}

	if len(arrMessage) > 0 {
		t.Errorf(strings.Join(arrMessage, "\n"))
	}
}
