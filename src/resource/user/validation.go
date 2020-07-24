package user

import (
	"fmt"
	"note/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserModelValidator struct {
	User struct {
		Username     string    `json:"username" binding:"required,min=8,max=255"`
		Password     string    `json:"password" binding:"required"`
	} `json: "user"`
	UserModel UserModel `json:"-"`
}

type UserLoginValidator struct {
	User struct {
		Username string `json:"username" binding:"required,min=8,max=255"`
		Password string `json:"password" binding:"required"`
	} `json:"user"`
	UserModel UserModel `json:"-"`
}

func (self *UserModelValidator) BindUser(c *gin.Context) error {
	err := common.Bind(c, self)
	fmt.Println(err)

	if err != nil {
		return err
	}

	self.UserModel.Username = self.User.Username
	self.UserModel.setPassword(self.User.Password)
	return nil
}

func (self *UserLoginValidator) BindLogin(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.UserModel.Username = self.User.Username
	self.UserModel.setPassword(self.User.Password)
	return nil
}

// IsExist is exist UserModel
func IsExist(user UserModel, db *gorm.DB) (UserModel, bool) {
	u := UserModel{}
	db.Where("username = ?", user.Username).First(&u)
	if u.Username == user.Username {
		return u, true
	}
	return u, false
}

//func FindUserById(id uuid.UUID, db *gorm.DB) (UserModel, bool) {
//	u := UserModel{}
//	db.Where("id = ?", id).First(&u)
//	if u.Username == user.Username {
//		return u, true
//	}
//	return u, false
//}
