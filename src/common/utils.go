package common

import (
	"note/config"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

func ResError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

func QueryList(resource interface{}, c *gin.Context) (interface{}, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	db.Find(resource)
	return resource, nil
}
