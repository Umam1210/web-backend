package models

type Bookmark struct {
	// ID        int `json:"id" gorm:"Bookmark_Id" `
	UserID    int     `json:"user_id"`
	User      User    `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ArtikelId int     `json:"artikel_id"`
	Artikel   Artikel `json:"artikel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type BookmarkResponse struct {
	UserID    int     `json:"user_id"`
	ArtikelId int     `json:"artikel_id"`
	Artikel   Artikel `json:"artikel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (BookmarkResponse) TableName() string {
	return "bookmarks"
}
