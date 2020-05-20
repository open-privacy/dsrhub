package model

import (
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    // mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // sqlite driver

	"github.com/avast/retry-go"
	"github.com/dsrhub/dsrhub/pkg/config"
	"go.uber.org/zap"
)

var (
	singletonDB   *gorm.DB
	singletonOnce sync.Once
)

// AutoMigrateTables stores the entity tables that we can auto migrate in gorm
var AutoMigrateTables = []interface{}{
	EncryptionScope{},
	Task{},
	Workflow{},
}

func connectDB() (db *gorm.DB, err error) {
	err = retry.Do(
		func() error {
			db, err = gorm.Open(config.ENV.DBDriver, config.ENV.DBConnectionStr)
			return err
		},
		retry.Attempts(config.ENV.DBConnectionRetryAttempts),
		retry.Delay(config.ENV.DBConnectionRetryDelay),
	)
	return db, err
}

// GetDB gets the db singleton
func GetDB() *gorm.DB {
	singletonOnce.Do(func() {
		db, err := connectDB()
		if err != nil {
			if config.ENV.DBConnectionDebug {
				config.Logger.Fatal("failed to connect to db", zap.Error(err))
			} else {
				config.Logger.Fatal("failed to connect to db")
			}
		}
		db.SetLogger(config.DBLogger)
		db.Debug().AutoMigrate(AutoMigrateTables...)
		singletonDB = db
	})

	return singletonDB
}

// NewSQLiteDB creates a new sqlite db
// useful for backup exports and unit tests
func NewSQLiteDB(filePath string) *gorm.DB {
	os.Remove(filePath)
	db, err := gorm.Open("sqlite3", filePath)
	if err != nil {
		config.Logger.Fatal("failed to connect to sqlite", zap.Error(err))
	}
	db.SetLogger(config.DBLogger)
	db.AutoMigrate(AutoMigrateTables...)
	return db
}

// NewTestDB creates a new test db
func NewTestDB() *gorm.DB {
	return NewSQLiteDB(":memory:")
}

// PopulateTestDB seeds the test db
func PopulateTestDB() *gorm.DB {
	testDB := NewTestDB()
	return testDB
}
