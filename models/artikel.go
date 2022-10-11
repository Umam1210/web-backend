package models

import "time"

type Artikel struct {
	ID     int       `json:"id" `
	Title  string    `json:"title" gorm:"type: varchar(255)"`
	Image  string    `json:"image" gorm:"type: varchar(255)"`
	Desc   string    `json:"desc" gorm:"type:varchar(255)"`
	Date   time.Time `json:"date"`
	User   User      `json:"user"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID int       `json:"user_id"`
}

type ArtikelResponse struct {
	ID    int    `json:"id" `
	Title string `json:"title"`
	Image string `json:"image"`
	Desc  string `json:"desc"`
}

func (ArtikelResponse) TableName() string {
	return "artikels"
}
