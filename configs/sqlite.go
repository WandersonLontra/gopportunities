package configs

import (
	"os"

	"github.com/WandersonLontra/gopportunities/internal/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("Database");
	dbPath := "./db/main.db"

	// Check if the database already exists
	_,err := os.Stat(dbPath);

	if os.IsNotExist(err){
		logger.Info("Database files not found, creating a new one ...")
		err = os.MkdirAll("./db",os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create database connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error on open sqlite database : %v", err)
		return nil, err
	}
	// Migrate the Schema
	err = db.AutoMigrate(&schema.Opening{})
	if err != nil {
		logger.Errorf("Error on migrate sqlite's schema : %v", err)
		return nil, err
	}

	// Return the DB pointer
	return db, nil
}