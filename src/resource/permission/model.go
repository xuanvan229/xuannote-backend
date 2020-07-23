package permission

import (
	"github.com/satori/uuid"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

// PermissionModel is a struct defiind Permission Table in DB
type PermissionModel struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Code string `json:"code"`
	Title string `json:"title"`
	Score int `json:"score"`
}

// BeforeCreate Model
func (permission *PermissionModel) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

// CreateTableUser Create Table if it doesn't
func CreateTableUser(db *gorm.DB) {
	check := db.HasTable(&PermissionModel{})
	if !check {
		db.CreateTable(&PermissionModel{})
	}
}
