package routes

import (
	"journey/handlers"
	"journey/pkg/middleware"
	"journey/pkg/mysql"
	"journey/repositories"

	"github.com/gorilla/mux"
)

func BookmarkRoutes(r *mux.Router) {
	bookmarkRepository := repositories.RepositoryBookmark(mysql.DB)
	h := handlers.HandlerBookmark(bookmarkRepository)

	r.HandleFunc("/bookmarks", h.FindBookmarks).Methods("GET")
	r.HandleFunc("/bookmark/{id}", h.GetBookmark).Methods("GET")
	r.HandleFunc("/bookmark", middleware.Auth(h.CreateBookmark)).Methods("POST")
	// r.HandleFunc("/bookmark/{id}", h.UpdateBookmark).Methods("PATCH")
	r.HandleFunc("/bookmark/{id}", h.DeleteBookmark).Methods("DELETE")

}
