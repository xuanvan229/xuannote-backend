package category

import (
	"note/common"
	//"note/config"
	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
)

type CategoriesModalValidate struct {
	Categories []struct {
		Title string `json:"title" binding:"required"`
		Image string `json:"image" binding:"required"`
	} `json:"categories"'`
	CategoriesModel []Category `json:"-"`
}

type CategoryModelValidator struct {
	Category struct {
		Title string `json:"title" binding:"required"`
		Image string `json:"image" binding:"required"`
	} `json: "category"`
	CategoryModel Category `json:"-"`
}

func NewCategoryModelValidator() CategoryModelValidator {
	return CategoryModelValidator{}
}

func NewCategoriesModelValidator() CategoriesModalValidate {
	return CategoriesModalValidate{}
}

func (self *CategoryModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.CategoryModel.Title = self.Category.Title
	self.CategoryModel.Image = self.Category.Image
	return nil
}

func (self *CategoriesModalValidate) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	for _, category := range self.Categories {
		nullUUID := uuid.NullUUID{}
		item := Category{
			Title:   category.Title,
			Image:   category.Image,
			OwnerID: nullUUID.UUID,
		}
		self.CategoriesModel = append(self.CategoriesModel, item)
	}
	return nil
}
