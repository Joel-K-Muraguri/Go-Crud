package seed

import (
	"log"

	"github.com/Joel-K-Muraguri/Go-Crud/api/models"
	"github.com/jinzhu/gorm"
)

var games = []models.Game{
	models.Game{
		Title : "God of War: Ragnarok",
	    Description : "Journey and encounter with norse Gods",
		Price: "KSH 8000",
		OS: "PS 5, PS4, PC , Xbox 360",
	},
	models.Game{
		Title : "Call of Duty Modern Warfare",
	    Description : "Ghost's story",
		Price: "KSH 7000",
		OS: "PS5, PS4, PC , Xbox 360",
	},
}

func Load(db *gorm.DB){

	err := db.Debug().DropTableIfExists(&models.Game{}).Error
	if err != nil {
		log.Fatalf("cannot drop table : %v", err)		
	}

	err = db.Debug().AutoMigrate(&models.Game{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table : %v", err)
		
	}


	for i, _ := range games {
		err = db.Debug().Model(&models.Game{}).Create(&games[i]).Error
		if err != nil {
			log.Fatalf("cannot seed games table : %v", err)		
		}
	}
}