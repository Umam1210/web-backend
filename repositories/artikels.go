package repositories

import (
	"journey/models"

	"gorm.io/gorm"
)

type ArtikelRepository interface {
	FindArtikels() ([]models.Artikel, error)
	GetArtikel(ID int) (models.Artikel, error)
	CreateArtikel(Artikel models.Artikel) (models.Artikel, error)
	UpdateArtikel(Artikel models.Artikel) (models.Artikel, error)
	DeleteArtikel(Artikel models.Artikel) (models.Artikel, error)
}

func RepositoryArtikel(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindArtikels() ([]models.Artikel, error) {
	var artikels []models.Artikel
	err := r.db.Preload("User").Find(&artikels).Error

	return artikels, err
}

func (r *repository) GetArtikel(ID int) (models.Artikel, error) {
	var artikel models.Artikel
	err := r.db.Preload("User").First(&artikel, ID).Error

	return artikel, err

}

func (r *repository) CreateArtikel(artikel models.Artikel) (models.Artikel, error) {
	err := r.db.Preload("User").Create(&artikel).Error

	return artikel, err
}

func (r *repository) UpdateArtikel(artikel models.Artikel) (models.Artikel, error) {
	err := r.db.Save(&artikel).Error

	return artikel, err
}

func (r *repository) DeleteArtikel(artikel models.Artikel) (models.Artikel, error) {
	err := r.db.Delete(&artikel).Error

	return artikel, err
}
