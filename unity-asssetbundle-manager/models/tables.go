package models

import (
	"github.com/jinzhu/gorm"
)

// AssetBundle info from Unity
type AssetBundle struct {
	gorm.Model

	Name      string `gorm:"unique_index"`
	Available bool
}
