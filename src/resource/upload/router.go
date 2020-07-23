package upload

import (
	"fmt"
	"net/http"
	"note/config"
	"note/utils/constant"
	"note/utils/md5"
	r "note/utils/response"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func UploadRegiter(router *gin.RouterGroup) {
	router.StaticFS("/upload/images", http.Dir(GetImageFullPath()))
	router.POST("/image", UploadImage)
}

func UploadImage(c *gin.Context) {
	app := r.Gin{C: c}
	_, header, err := c.Request.FormFile("image")
	userIDString := c.Request.Header.Get("_userID")
	db, err := config.Connect()
	defer db.Close()
	// userID, err := strconv.ParseUint(userIDString, 10, 64)
	userID, err := uuid.FromString(userIDString)
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
	}

	if header == nil {
		fmt.Println("hello")
		app.Response(http.StatusBadRequest, constant.ERROR, nil, nil)
		return
	}
	imageName := GetImageName(header.Filename)
	fullPath := GetImageFullPath()
	src := fullPath + imageName

	if err := c.SaveUploadedFile(header, src); err != nil {
		fmt.Println("eror", err)
		app.Response(http.StatusInternalServerError, constant.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil, nil)
		return
	}

	CreateTableImageModel(db)
	newImageUpload := ImageUploadModel{}
	newImageUpload.Filename = imageName
	newImageUpload.FileURL = GetImageFullUrl(imageName)
	newImageUpload.UserID = userID
	db.Create(&newImageUpload)
	app.Response(http.StatusOK, constant.SUCCESS, map[string]string{
		"image": config.PrefixUrl + GetImageFullUrl(imageName),
	}, nil)

}

func GetImageFullUrl(name string) string {
	return "/api/" + config.ImageSavePath + "/" + name
}

func ViewImageRegiter(router *gin.RouterGroup) {
	router.StaticFS("/upload/images", http.Dir(GetImageFullPath()))
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = md5.EncodeMD5(fileName)

	return fileName + ext
}

func CreateTableImageModel(db *gorm.DB) {
	check := db.HasTable(&ImageUploadModel{})
	if !check {
		db.CreateTable(&ImageUploadModel{})
	}
}

// GetImageFullPath get full save path
func GetImageFullPath() string {
	return config.RuntimeRootPath + "/" + config.ImageSavePath + "/"
}
