package user

import (
	// "github.com/satori/go.uuid"
	"fmt"
	"net/http"
	"note/config"
	"note/utils/constant"
	r "note/utils/response"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var null_data = map[string]string{}

// UserRegister router for user
func UserRegister(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/logout", Logout)
}

// Register control
func Register(c *gin.Context) {
	app := r.Gin{C: c}
	userModelValidator := NewUserModelValidator()

	if err := userModelValidator.BindUser(c); err != nil {
		app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
		return
	}
	db, err := config.Connect()
	CreateTableUser(db)
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
		return
	}

	fmt.Println("userModelValidator", userModelValidator)
	_, isExist := IsExist(userModelValidator.UserModel, db)

	if isExist {
		app.Response(http.StatusInternalServerError, constant.USER_NAME_IS_EXIST, nil, nil)
		return
	}

	db.Create(&userModelValidator.UserModel)
	app.Response(http.StatusOK, constant.SUCCESS, userModelValidator.UserModel, nil)
	return
}

// Logout user
func Logout(c *gin.Context) {
	app := r.Gin{C: c}
	c.SetCookie("_token", "", 0, "/", "", false, true)
	app.Response(http.StatusOK, constant.SUCCESS, map[string]string{"logout": "ok"}, nil)
	return
}

// Login user
func Login(c *gin.Context) {
	app := r.Gin{C: c}
	userLoginValidator := NewUserLoginValidator()

	if err := userLoginValidator.BindLogin(c); err != nil {
		app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
		return
	}

	fmt.Print(userLoginValidator)

	db, err := config.Connect()
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.COULD_NOT_CONNECT_TO_DATABASE, nil, nil)
		return
	}

	user, isExist := IsExist(userLoginValidator.UserModel, db)

	if !isExist {
		app.Response(http.StatusInternalServerError, constant.USER_DOES_NOT_EXIST, nil, nil)
		return
	}

	if comparePasswords(user.Password, []byte(userLoginValidator.User.Password)) {
		token, err := createToken(user)
		if err != nil {
			app.Response(http.StatusInternalServerError, constant.WRONG_PASSWORD, nil, nil)
			return
		}
		c.SetCookie("_token", token, 3600, "/", "", false, true)
		c.SetCookie("_userID", user.ID.String(), 3600, "/", "", false, true)
		data := map[string]interface{}{
			"token":      token,
			"userID":     user.ID.String(),
		}
		app.Response(200, constant.SUCCESS, data, map[string]int{})
		return
	}

	app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
	return
}
