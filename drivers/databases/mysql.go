package mysql_driver

import (
	"fmt"
	"log"
	"os"
	"time"

	"weatherit/drivers/databases/events"
	"weatherit/drivers/databases/interests"
	"weatherit/drivers/databases/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func (config *ConfigDB) InitialDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Database)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&events.Event{},
		&events.EventChecklist{},
		&users.User{},
		&interests.Interest{},
	)

	// init interests
	interests := []interests.Interest{
		{ID: 1, Name: "Lawn & Garden"},
		{ID: 2, Name: "Fashion"},
		{ID: 3, Name: "Animals"},
		{ID: 4, Name: "Photography"},
		{ID: 5, Name: "Art"},
		{ID: 6, Name: "Home Decor"},
		{ID: 7, Name: "DIY & Crafts"},
		{ID: 8, Name: "Education"},}
	db.Create(&interests)
}
