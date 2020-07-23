package category

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
	"note/resource/user"
)

type Category struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Title     string         `json:"title"`
	Image     string         `json:"image"`
	UserOwner user.UserModel `json:"-"`
	OwnerID   uuid.UUID      `json:"owner_id"`
}

type CategoryDelete struct {
	ID string `json:"id" binding:"required"`
}

type CategoryUpdate struct {
	ID    string `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Image string `json:"image" binding:"required"`
}

// BeforeCreate Model
func (wallet *Category) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()
	return scope.SetColumn("ID", id)
}

// CreateTableUser Create Table if it doesn't
func CreateTableCategory(db *gorm.DB) {
	check := db.HasTable(&Category{})
	if !check {
		db.CreateTable(&Category{})
	}
}

func InsertMany(db *gorm.DB, categories []Category) {
	for _, category := range categories {
		db.Create(&category)
	}
}
