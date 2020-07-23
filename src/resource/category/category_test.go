package category

import (
	"math/rand"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"

	//"note/config"
	"fmt"
	"testing"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func stubUsers(b *testing.B) (categories []*Category) {
	for i := 0; i < b.N; i++ {
		nullUUID := uuid.NullUUID{}
		category := &Category{
			Title:   genRandomString(100),
			Image:   genRandomString(100),
			OwnerID: nullUUID.UUID,
		}
		categories = append(categories, category)
	}

	return categories
}

func BenchmarkOrmCreate(b *testing.B) {
	//db, err := config.Connect()
	db, err := gorm.Open("postgres", "host=128.199.152.226 port=5432 user=postgres dbname=money password=conchoancut1")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	categories := stubUsers(b)

	for _, category := range categories {
		fmt.Println("database,", category)
		db.Create(&category)
	}
}

func genRandomString(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
