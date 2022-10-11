package artikelsdto

import "time"

type CreateArtikelRequest struct {
	Title   string    `json:"title" form:"title" validate:"required"`
	Image   string    `json:"image" form:"image" validate:"required"`
	Desc    string    `json:"desc" form:"desc" validate:"required"`
	User_Id int       `json:"user_id"`
	Date    time.Time `json:"date"`
}

type UpdateArtikelRequest struct {
	Title string `json:"title" form:"title"`
	Image string `json:"thumbnail" form:"thumbnail"`
	Desc  string `json:"desc" form:"desc"`
}
