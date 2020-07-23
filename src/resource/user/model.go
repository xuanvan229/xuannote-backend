package user

import (
	"log"
	"note/config"
	"note/resource/permission"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt    time.Time `json:"create_at"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Password     string    `json:"-"`
	Permission   permission.PermissionModel
	PermissionID uuid.UUID `json:"-"`
}

type JwtCustomClaims struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	jwt.StandardClaims
}

func (user *UserModel) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

func (user *UserModel) setPassword(pw string) error {
	pwd := []byte(pw)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	//fmt.Println(byteHash,plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func createToken(user UserModel) (string, error) {
	claims := &JwtCustomClaims{
		user.Username,
		user.ID.String(),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.AppSetting.JwtSecret))

	if err != nil {
		return t, err
	}

	return t, nil
}

func NewUserModelValidator() UserModelValidator {
	return UserModelValidator{}
}

func NewUserLoginValidator() UserLoginValidator {
	return UserLoginValidator{}
}

func CreateTableUser(db *gorm.DB) {
	check := db.HasTable(&UserModel{})
	if !check {
		db.CreateTable(&UserModel{})
	}
}
