package note

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type NoteModel struct {
	ID       uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Title    string         `json:"title"`
	Content  string         `json:"content"`
	Category string			`json:"category"`
}

func NewNoteModelValidator() NoteModelValidator {
	return NoteModelValidator{}
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *NoteModel) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	fmt.Println("uuid", uuid)
	return scope.SetColumn("ID", uuid)
}

func CreateTableUser(db *gorm.DB) {
	check := db.HasTable(&NoteModel{})
	if !check {
		db.CreateTable(&NoteModel{})
	}
}
