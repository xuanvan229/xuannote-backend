package note

import (
	// "errors"
	"note/common"

	"github.com/gin-gonic/gin"
)

type NoteModelValidator struct {
	Note struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		Category string `json:"category"`

	} `json:"note"`
	NoteModel NoteModel `json:"-"`
}

func (self *NoteModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)

	if err != nil {
		return err
	}
	self.NoteModel.Title = self.Note.Title
	self.NoteModel.Content = self.Note.Content
	self.NoteModel.Category = self.Note.Category
	return nil
}

