package permission

import (
	"github.com/gin-gonic/gin"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"note/config"
	"note/utils/constant"
	r "note/utils/response"
)

func PermissionRegister(router *gin.RouterGroup) {
	router.POST("/permission", PutRegister)
	router.GET("/permission", GetRegister)
}

func PutRegister(c *gin.Context) {
	app := r.Gin{C: c}
	permission := NewPermissionModelValidator()
	if err := permission.Bind(c); err != nil {
		log.Fatal("error", err)
		app.Response(http.StatusInternalServerError, constant.MISSING_SOME_FIELD, nil, nil)
		return
	}

	db, err := config.Connect()
	if err != nil {

		app.Response(http.StatusInternalServerError, constant.COULD_NOT_CONNECT_TO_DATABASE, nil, nil)
		return
	}
	CreateTableUser(db)

	_, isExist := IsExist(permission.PermissionModel, db)

	if isExist {
		app.Response(http.StatusInternalServerError, constant.PERMISSION_ALREADY_EXIST, nil, nil)
		return
	}

	db.Create(&permission.PermissionModel)
	app.Response(http.StatusOK, constant.SUCCESS, permission.PermissionModel, nil)
	return

}

func GetRegister(c *gin.Context) {

}
