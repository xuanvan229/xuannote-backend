package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"note/common"
	"note/config"
	"note/resource/user"
	"note/utils/constant"
	r "note/utils/response"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var ctx = context.Background()

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3002")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods:", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers: Content-Type", "*")
		next.ServeHTTP(w, r)
		return
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, _userID")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		app := r.Gin{C: c}
		token := c.Request.Header.Get("Authorization")
		userIdToken := c.Request.Header.Get("_userID")
		fmt.Println("token", token)
		if len(token) == 0 {
			fmt.Println("error author")
			app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
			c.Abort()
			return
		}
		//userIdToken, err := c.Cookie("_userID")
		if len(userIdToken) == 0 {
			fmt.Println("error author")
			app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
			c.Abort()
			return
		}

		userToken := user.JwtCustomClaims{}
		_, err := jwt.ParseWithClaims(token, &userToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		fmt.Println("userToken token, ", userToken)

		if err != nil {
			c.JSON(404, map[string]string{"status": "not ok 1234"})
			c.Abort()
			return
		}
		userInfo := user.UserModel{}

		if userToken.ID != userIdToken {
			fmt.Println("the error is here")
			app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
			c.Abort()
			return
		}

		userID, err := uuid.FromString(userToken.ID)
		if err != nil {
			c.JSON(404, map[string]string{"status": "not ok 1234"})
			c.Abort()
			return
		}

		userInfo.Username = userToken.Username
		userInfo.ID = userID

		db, err := config.Connect()
		if err != nil {
			c.JSON(503, common.ResError("user", err))
			c.Abort()
			return
		}

		defer db.Close()

		_, isExist := user.IsExist(userInfo, db)
		fmt.Println("the error is here", isExist)

		if !isExist {
			c.JSON(503, common.ResError("user", errors.New("Use does not exist")))
			c.Abort()
			return
		} else {
			c.Set("username", userInfo.Username)
			c.Next()
		}
		return
	}
}

func CheckLogin(c *gin.Context) {
	app := r.Gin{C: c}
	cookie, err := c.Cookie("_token")
	if err != nil {
		fmt.Println(err)
		app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
		//c.JSON(404, map[string]string{"status": "not ok"})
		return
	}

	userToken := user.JwtCustomClaims{}
	_, err = jwt.ParseWithClaims(cookie, &userToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
		return
	}
	userID, err := uuid.FromString(userToken.ID)
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
		return
	}

	userInfo := user.UserModel{}

	userInfo.Username = userToken.Username
	userInfo.ID = userID

	db, err := config.Connect()
	if err != nil {
		app.Response(http.StatusInternalServerError, constant.ERROR, nil, nil)
		return
	}
	defer db.Close()
	userInDB, isExist := user.IsExist(userInfo, db)

	fmt.Println("userInDB", userInDB)
	if !isExist {
		app.Response(http.StatusInternalServerError, constant.USER_DOES_NOT_EXIST, nil, nil)
		return
	}
	user_info := map[string]interface{}{
		"username":   userInDB.Username,
	}
	app.Response(http.StatusOK, constant.SUCCESS, user_info, nil)
	return
}
