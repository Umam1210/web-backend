package bookmarksdto

import (
	artikelsdto "journey/dto/artikel"
	"journey/models"
)

type BookmarkResponse struct {
	ID         int
	UserID     int                         `json:"userId"`
	Artikel_Id int                         `json:"artikel_id"`
	User       models.User                 `json:"user"`
	Artikel    artikelsdto.ArtikelResponse `json:"artikel"`
}
