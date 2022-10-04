package routes

import (
	"journey/handlers"
	"journey/pkg/middleware"
	"journey/pkg/mysql"
	"journey/repositories"

	"github.com/gorilla/mux"
)

func ArtikelRoutes(r *mux.Router) {
	artikelRepository := repositories.RepositoryArtikel(mysql.DB)
	h := handlers.HandlerArtikel(artikelRepository)

	r.HandleFunc("/artikels", h.FindArtikels).Methods("GET")
	r.HandleFunc("/artikel/{id}", h.GetArtikel).Methods("GET")
	r.HandleFunc("/artikel", middleware.Auth(middleware.UploadFile(h.CreateArtikel))).Methods("POST")
	r.HandleFunc("/artikel/{id}", h.UpdateArtikel).Methods("PATCH")
	r.HandleFunc("/artikel/{id}", h.DeleteArtikel).Methods("DELETE")

}
