package upload

import (
	"note/resource/user"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// ImageUploadModel struct of image upload
type ImageUploadModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time      `json:"create_at"`
	Filename  string         `json:"file_name"`
	User      user.UserModel `json:"-"`
	UserID    uuid.UUID      `gorm:"type:uuid;" json:"user_id"`
	FileURL   string         `json:"file_url"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *ImageUploadModel) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}
