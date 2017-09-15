package sqlite3

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //gorm dialect
	"private/gts.com/asssetbundle-manager/models"
)

// Manager manage connection
type Manager struct {
	Error error
	DB    *gorm.DB
}

// Open connect to sqlite3 database
func Open(path string) *Manager {
	db, err := gorm.Open("sqlite3", path)

	return &Manager{
		DB:    db,
		Error: err,
	}
}

// Migrate table schema
func (m *Manager) Migrate() *Manager {
	m.DB.AutoMigrate(&models.AssetBundle{})

	return m
}

// Close connection
func (m *Manager) Close() {
	m.Error = m.DB.Close()
}
