package handlers

import (
	"errors"
	"go-xsis-movie/helpers"
	"go-xsis-movie/models/db"
	"go-xsis-movie/models/entities"
	"time"

	"gorm.io/gorm"
)

type H map[string]interface{}
type Movie entities.Movie

func (h *Movie) H_AddMovie() (H, error) {
	datum := Movie{}
	datum.Title = h.Title
	datum.Description = h.Description
	datum.Rating = h.Rating
	datum.Image = h.Image

	err := db.GetDB().Debug().Create(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	rMsg := H{}
	rMsg["id"] = datum.Id
	rMsg["title"] = datum.Title
	rMsg["description"] = datum.Description
	rMsg["rating"] = datum.Rating
	rMsg["image"] = datum.Image
	rMsg["created_at"] = datum.CreatedAt
	rMsg["updated_at"] = datum.UpdatedAt

	return rMsg, nil
}

func H_ReadAllMovie() (*[]Movie, error) {
	var datum []Movie
	err := db.GetDB().Debug().Find(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		return nil, err
	}
	return &datum, nil
}

func H_ReadSingleMovieId(id string) (*Movie, error) {
	var datum Movie
	err := db.GetDB().Debug().Where("id = ?", id).Find(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		return nil, errors.New("Movie with ID not found")
	}
	return &datum, nil
}

func (h *Movie) H_UpdateMovieId(id string) (*Movie, error) {
	var existingMovie Movie
	err := db.GetDB().Debug().Model(Movie{}).Where("id = ?", id).First(&existingMovie).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Movie with ID not found")
		}
		helpers.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	existingMovie.Title = h.Title
	existingMovie.Description = h.Description
	existingMovie.Rating = h.Rating
	existingMovie.Image = h.Image
	existingMovie.UpdatedAt = time.Now()

	err = db.GetDB().Debug().Save(&existingMovie).Error
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	return &existingMovie, nil
}

func H_DeleteMovie(id string) (string, error) {
	var movie Movie
	err := db.GetDB().Debug().Model(Movie{}).Where("id = ?", id).First(&movie).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("Movie with ID not found")
		}
		helpers.Logger("error", "In Server: "+err.Error())
		return "", err
	}

	err = db.GetDB().Debug().Delete(&movie).Error
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		return "", err
	}

	return "Delete Movie From Database", err
}
