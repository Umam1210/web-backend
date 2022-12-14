package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	artikelsdto "journey/dto/artikel"
	dto "journey/dto/result"
	"journey/models"
	"journey/repositories"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerArtikel struct {
	ArtikelRepository repositories.ArtikelRepository
}

var PathFile = os.Getenv("PATH_FILE")

func HandlerArtikel(ArtikelRepository repositories.ArtikelRepository) *handlerArtikel {
	return &handlerArtikel{ArtikelRepository}
}

func (h *handlerArtikel) FindArtikels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	artikels, err := h.ArtikelRepository.FindArtikels()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	// for i, p := range artikels {
	// 	artikels[i].Image = os.Getenv("PATH_FILE") + p.Image
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: artikels}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerArtikel) FindArtikelsbyUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user_id, _ := strconv.Atoi(mux.Vars(r)["user_id"])

	artikels, err := h.ArtikelRepository.FindArtikelsbyUserId(user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	// for i, p := range artikels {
	// 	artikels[i].Image = os.Getenv("PATH_FILE") + p.Image
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: artikels}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArtikel) GetArtikel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var artikel models.Artikel

	artikel, err := h.ArtikelRepository.GetArtikel(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// artikel.Image = os.Getenv("PATH_FILE") + artikel.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseArtikel(artikel)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArtikel) CreateArtikel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	fmt.Println(userId)
	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)
	Date := time.Now()

	request := artikelsdto.CreateArtikelRequest{
		Title:   r.FormValue("title"),
		Desc:    r.FormValue("desc"),
		Image:   filepath,
		User_Id: userId,
		Date:    Date,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "journey"})

	if err != nil {
		fmt.Println(err.Error())
	}

	artikel := models.Artikel{
		Title:  request.Title,
		Image:  resp.SecureURL,
		Desc:   request.Desc,
		UserID: request.User_Id,
		Date:   request.Date,
	}

	data, err := h.ArtikelRepository.CreateArtikel(artikel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	artikel, _ = h.ArtikelRepository.GetArtikel(data.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseArtikel(artikel)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArtikel) UpdateArtikel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(artikelsdto.UpdateArtikelRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	artikel, err := h.ArtikelRepository.GetArtikel(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		artikel.Title = request.Title
	}

	if request.Image != "" {
		artikel.Image = request.Image
	}

	if request.Desc != "" {
		artikel.Desc = request.Desc
	}

	data, err := h.ArtikelRepository.UpdateArtikel(artikel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseArtikel(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArtikel) DeleteArtikel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	artikel, err := h.ArtikelRepository.GetArtikel(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ArtikelRepository.DeleteArtikel(artikel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseArtikel(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseArtikel(u models.Artikel) artikelsdto.ArtikelResponse {
	return artikelsdto.ArtikelResponse{
		ID:    u.ID,
		Title: u.Title,
		Image: u.Image,
		Desc:  u.Desc,
		User:  u.User,
		Date:  u.Date,
	}
}
