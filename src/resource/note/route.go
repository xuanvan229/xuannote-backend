package note

import (
	"fmt"
	"net/http"
	"note/common"
	"note/config"
	"note/utils/constant"
	r "note/utils/response"

	"github.com/gin-gonic/gin"
)

func RegisterRouteNoToken(router *gin.RouterGroup) {
	router.GET("/note", GetNote)
	router.GET("/note/:id", GetItem)
}

func RegisterRoute(router *gin.RouterGroup) {
	router.POST("/note", PostNote)

}

func GetNote(c *gin.Context) {
	app := r.Gin{C: c}
	listnote, err := common.QueryList(&[]NoteModel{}, c)
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
		return
	}
	app.Response(http.StatusOK, constant.SUCCESS, listnote, nil)
	return
}

func GetItem(c *gin.Context) {
	app := r.Gin{C: c}
	id := c.Param("id")
	db, err := config.Connect()
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
		return
	}
	defer db.Close()
	note := NoteModel{}
	db.Where("ID = ?", id).First(&note)
	app.Response(http.StatusOK, constant.SUCCESS, note, nil)
	return
}

func PostNote(c *gin.Context) {
	app := r.Gin{C: c}
	noteModelValidator := NewNoteModelValidator()
	if err := noteModelValidator.Bind(c); err != nil {
		app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
		return
	}
	fmt.Println("fmt",noteModelValidator )
	db, err := config.Connect()
	CreateTableUser(db)
	defer db.Close()
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
		return
	}
	db.Create(&noteModelValidator.NoteModel)
	app.Response(http.StatusOK, constant.SUCCESS, noteModelValidator.NoteModel, nil)
	return
}
