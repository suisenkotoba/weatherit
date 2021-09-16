package mysql_driver

import (
	"fmt"
	"log"
	"os"
	"time"

	"weatherit/drivers/databases/events"
	"weatherit/drivers/databases/alterplans"
	"weatherit/drivers/databases/activities"
	"weatherit/drivers/databases/interests"
	"weatherit/drivers/databases/users"
	userInterests "weatherit/drivers/databases/user_interests"

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
		&activities.Activity{},
		&alterplans.AlterPlan{},
		&users.User{},
		&interests.Interest{},
		&userInterests.UserInterest{},
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
		{ID: 8, Name: "Education"},
		{ID: 9, Name: "Movies"},
	}

	// init activities
	activities := []activities.Activity{
		{ID: 1, Name: "Gardening", IsOutdoor: true, RecommendedWeather: "Clear", InterestID: 1},
		{ID: 2, Name: "Window Shopping", IsOutdoor: true, RecommendedWeather: "Clear", InterestID: 2},
		{ID: 3, Name: "Knitting", IsOutdoor: false, RecommendedWeather: "Thunderstorm", InterestID: 7},
		{ID: 4, Name: "Resin Molding", IsOutdoor: false, RecommendedWeather: "Clear", InterestID: 5},
		{ID: 5, Name: "Binge Watching Last Season Anime", IsOutdoor: false, RecommendedWeather: "Rain", InterestID: 9},
	}
	db.Create(&interests)
	db.Create(&activities)
}
