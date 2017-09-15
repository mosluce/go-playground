package models

import "time"

// AssetBundle info from Unity
type AssetBundle struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Name      string `gorm:"unique_index" json:"name"`
	Available bool   `json:"available"`
}
