package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"note/config"
	"note/utils/constant"
	r "note/utils/response"
)

func CategoryRegister(router *gin.RouterGroup) {
	router.POST("/categories", AddCategories)
	router.POST("/category", AddCategory)
	router.GET("/categories", GetCategory)
	router.PUT("/category", UpdateCategory)
	router.DELETE("/category", DelCategory)
}

func AddCategories(c *gin.Context) {
	app := r.Gin{C: c}

	categoriesValidate := NewCategoriesModelValidator()
	if err := categoriesValidate.Bind(c); err != nil {
		app.Response(http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	db, err := config.Connect()
	defer db.Close()
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.COULD_NOT_CONNECT_TO_DATABASE, nil, nil)
		return
	}

	InsertMany(db, categoriesValidate.CategoriesModel)
	app.Response(http.StatusOK, constant.SUCCESS, categoriesValidate.CategoriesModel, nil)
	return
}

func AddCategory(c *gin.Context) {
	app := r.Gin{C: c}
	userIDString := c.Request.Header.Get("_userID")
	userID, err := uuid.FromString(userIDString)

	categoryValidator := NewCategoryModelValidator()
	if err := categoryValidator.Bind(c); err != nil {
		app.Response(http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	db, err := config.Connect()
	defer db.Close()
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.COULD_NOT_CONNECT_TO_DATABASE, nil, nil)
		return
	}

	CreateTableCategory(db)
	categoryValidator.CategoryModel.OwnerID = userID
	db.Create(&categoryValidator.CategoryModel)
	app.Response(http.StatusOK, constant.SUCCESS, categoryValidator.CategoryModel, nil)
	return
}

func GetCategory(c *gin.Context) {
	app := r.Gin{C: c}
	userIDString := c.Request.Header.Get("_userID")
	db, err := config.Connect()
	defer db.Close()
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.COULD_NOT_CONNECT_TO_DATABASE, nil, nil)
		return
	}
	categories := FindCategories(db, userIDString)
	app.Response(http.StatusOK, constant.SUCCESS, categories, nil)
	return
}

func UpdateCategory(c *gin.Context) {
	app := r.Gin{C: c}
	userIDString := c.Request.Header.Get("_userID")
	db, err := config.Connect()
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.COULD_NOT_CONNECT_TO_DATABASE, nil, nil)
		return
	}

	var json CategoryUpdate
	var category Category
	if err := c.ShouldBindJSON(&json); err != nil {
		app.Response(http.StatusBadRequest, err.Error(), nil, nil)
		return
	}
	if err := db.Where("id = ? AND owner_id = ?", json.ID, userIDString).First(&category).Error; err != nil {
		app.Response(http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	category.Title = json.Title
	category.Image = json.Image
	if err := db.Save(&category).Error; err != nil {
		app.Response(http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	app.Response(http.StatusOK, constant.SUCCESS, map[string]string{
		"message": "success",
	}, nil)
	return

}

func DelCategory(c *gin.Context) {
	app := r.Gin{C: c}
	userIDString := c.Request.Header.Get("_userID")
	db, err := config.Connect()
	defer db.Close()
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.COULD_NOT_CONNECT_TO_DATABASE, nil, nil)
		return
	}

	var json CategoryDelete
	if err := c.ShouldBindJSON(&json); err != nil {
		app.Response(http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	category, err := FindCategory(json.ID, userIDString, db)

	if err != nil {
		app.Response(http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	if err := db.Delete(&category).Error; err != nil {
		app.Response(http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	app.Response(http.StatusOK, constant.SUCCESS, map[string]string{
		"message": "success",
	}, nil)
	return
}

func FindCategory(id string, userID string, db *gorm.DB) (Category, error) {
	var category Category
	if err := db.Where("id = ? AND owner_id = ?", id, userID).Find(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func FindCategories(db *gorm.DB, id string) []Category {
	categories := []Category{}
	nullUUId := uuid.NullUUID{}
	fmt.Println("null uuid", nullUUId.UUID)
	db.Where("owner_id = ? OR owner_id = ?", id, nullUUId.UUID).Find(&categories)
	return categories
}
