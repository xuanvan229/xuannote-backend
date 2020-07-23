package permission

import (
	"note/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

type PermissionModelValidator struct {
	Permission struct {
		Code  string `json:"code" binding:"required,min=2,max=255"`
		Title string `json:"title" binding:"required"`
		Score int    `json:"score" binding:"required"`
	} `json: "permission"`
	PermissionModel PermissionModel `json:"-"`
}

// NewPermissionModelValidator new Permission Validator
func NewPermissionModelValidator() PermissionModelValidator {
	return PermissionModelValidator{}
}

func (self *PermissionModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.PermissionModel.Code = self.Permission.Code
	self.PermissionModel.Title = self.Permission.Title
	self.PermissionModel.Score = self.Permission.Score
	return nil
}

func IsExist(permission PermissionModel, db *gorm.DB) (PermissionModel, bool) {
	p := PermissionModel{}
	db.Where("code = ?", permission.Code).First(&p)
	if p.Code == permission.Code {
		return p, true
	}

	return p, false
}
