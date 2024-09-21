package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config holds the database configuration settings
type PostgresQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
	TimeZone string `json:"timezone"`
}

// Service holds the database connection and current table context
type Service struct {
	DB *gorm.DB
}

// NewService creates a new database service
func (cfg PostgresQL) NewService() (*Service, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode, cfg.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Service{DB: db}, nil
}

// Create inserts a new record into the current table
func (s *Service) Create(record interface{}) error {
	if err := s.DB.Create(record).Error; err != nil {
		return fmt.Errorf("failed to create record: %w", err)
	}
	return nil
}

// gets a record from the table
func (s *Service) GetRecord(recID string) (interface{}, bool) {
	return s.DB.Get(recID)
}
