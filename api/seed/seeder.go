package seed

import (
	"log"

	"github.com/loeken/rest-blog/api/models"
	"gorm.io/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "loeken",
		Email:    "loeken@internetz.me",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther King",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:    "Title 1",
		Content:  "Hello world 1",
		AuthorID: 1,
	},
	models.Post{
		Title:    "Title 2",
		Content:  "Hello world 2",
		AuthorID: 1,
	},
}

func Load(db *gorm.DB) {
	db.Migrator().DropTable(&models.Post{}, &models.User{})
	err := db.AutoMigrate(&models.User{}, &models.Post{})
	//err := db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	/* @todo remove if not needed
	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	*/
	//db.Debug().Model(&models.User{}).Delete("*")
	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
