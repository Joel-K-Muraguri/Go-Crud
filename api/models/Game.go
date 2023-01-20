package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Game struct {
	ID          uint      `gorm:"primaryKey;auto_increment" json:"id"`
	Title        string    `gorm:"size:255;not null;unique" json:"title"`
	Description string    `gorm:"size:255;not null;" json:"description"`
	Price       string    `gorm:"size:255;not null;" json:"price"`
	OS          string    `gorm:"size:255;not null;" json:"os"`
	Created_At  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Update_At   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (g *Game) Prepare() {
	g.ID = 0
	g.Title = html.EscapeString(strings.TrimSpace(g.Name))
	g.Description = html.EscapeString(strings.TrimSpace(g.Description))
	g.Price = html.EscapeString(strings.TrimSpace(g.Price))
	g.OS = html.EscapeString(strings.TrimSpace(g.OS))
	g.Created_At = time.Now()
	g.Update_At = time.Now()

}

func (g *Game) Validate() error {

	if g.Title == "" {
		return errors.New("REQUIRED NAME")

	}
	if g.Description == "" {
		return errors.New("REQUIRED DESCRIPTION")

	}
	if g.Price == "" {
		return errors.New("REQUIRED PRICE")

	}
	if g.OS == "" {
		return errors.New("REQUIRED OS")

	}
	return nil

}

func (g *Game) SaveGame(db *gorm.DB) (*Game, error) {

	var err error
	err = db.Debug().Create(&g).Error
	if err != nil {
		return &Game{}, err
	}
	return g, nil
}

func (g *Game) FindAllGames(db *gorm.DB) (*[]Game, error) {
	var err error
	games := []Game{}
	err = db.Debug().Model(&Game{}).Limit(100).Find(&games).Error
	if err != nil {
		return &[]Game{}, err
	}
	return &games, err
}

func (g *Game) FindGameByID(db *gorm.DB, uid uint32) (*Game, error) {
	var err error
	err = db.Debug().Model(Game{}).Where("id = ?", uid).Take(&g).Error
	if err != nil {
		return &Game{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Game{}, errors.New("User Not Found")
	}
	return g, err
}

func (g *Game) UpdateAGame(db *gorm.DB, uid uint32) (*Game, error) {

	var err error
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Game{}).Where("id = ?", uid).Take(&Game{}).UpdateColumns(
		map[string]interface{}{
			"name":             g.Title,
			"description":      g.Description,
			"price":            g.Price,
			"operating_system": g.OS,
			"updated_at":       time.Now(),
		},
	)
	if db.Error != nil {
		return &Game{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&Game{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Game{}, err
	}
	return g, nil
}

func (g *Game) DeleteAGame(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Game{}).Where("id = ?", uid).Take(&Game{}).Delete(&Game{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
