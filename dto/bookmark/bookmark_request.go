package bookmarksdto

type CreateBookmarkRequest struct {
	// ID         int
	UserID     int `json:"userId"`
	Artikel_Id int `json:"artikel_id"`
}
