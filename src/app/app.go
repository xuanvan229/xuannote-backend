package app

import (
	// "github.com/gorilla/mux"
	// "encoding/json"
	"fmt"

	"note/config"
	"note/middleware"

	//"note/resource/page"

	"note/resource/category"
	"note/resource/note"
	"note/resource/permission"
	"note/resource/upload"
	"note/resource/user"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	// DB *config.PostGresConfig
}

func (a *App) Initialize() {
	config.Setup()
	a.Router = gin.New()
	a.setRouters()
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middle ware")
	}
}

func AuthRequired1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middle ware 1")
	}
}

func (app *App) setRouters() {
	store, _ := sessions.NewRedisStore(config.RedisSetting.Size, config.RedisSetting.NetWork, config.RedisSetting.Host, config.RedisSetting.Password, []byte(config.RedisSetting.KeyPairs))
	app.Router.Use(sessions.Sessions("session", store))
	app.Router.Use(middleware.CORSMiddleware())

	app.Router.Use(gin.Logger())
	app.Router.Use(gin.Recovery())
	api := app.Router.Group("/api")
	upload.ViewImageRegiter(api)
	note.RegisterRouteNoToken(api)
	api.GET("/hi", middleware.CheckLogin)
	obj := api.Group("/obj")
	obj.Use(middleware.AuthorizationMiddleware())
	note.RegisterRoute(obj)
	upload.UploadRegiter(obj)
	category.CategoryRegister(obj)
	u := api.Group("/user")
	u.Use(AuthRequired1())
	user.UserRegister(u)
	permission.PermissionRegister(obj)
}
