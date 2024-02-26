package repository

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"todo.app/utils"
)

func InitializeDatabase() gorm.DB {
	var connectionString = utils.GetEnvironmentVariable("DATABASE_CONNECTION_STRING")

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	db.Logger = logger.Default.LogMode(logger.Info)

	if err != nil {
		log.Panic().Msg("Failed to establish connection to postgres database")
	}

	err = db.AutoMigrate(&Todo{})

	if err != nil {
		log.Panic().Msg("Failed to migrate database schema")
	}

	log.Info().Msg("Data source migrated successfully")

	return *db
}
