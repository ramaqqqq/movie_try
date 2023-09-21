package controllers

import (
	"encoding/json"
	"go-xsis-movie/handlers"
	"go-xsis-movie/helpers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func C_AddMovie(w http.ResponseWriter, r *http.Request) {
	datum := handlers.Movie{}
	err := json.NewDecoder(r.Body).Decode(&datum)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	result, err := datum.H_AddMovie()
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(result)
	helpers.Logger("info", "successfully, created movie: "+string(logger))
	helpers.Response(w, http.StatusCreated, rMsg)
}

func C_GetAllMovie(w http.ResponseWriter, r *http.Request) {
	result, err := handlers.H_ReadAllMovie()
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	helpers.Logger("info", "read all movies")
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_GetSingleMovieId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idMovie := params["ID"]

	result, err := handlers.H_ReadSingleMovieId(idMovie)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	helpers.Logger("info", "readOne all movies")
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_UpdateSingleMovieId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idMovie := params["ID"]

	datum := handlers.Movie{}
	err := json.NewDecoder(r.Body).Decode(&datum)

	result, err := datum.H_UpdateMovieId(idMovie)
	if err != nil {

		if gorm.IsRecordNotFoundError(err) {
			helpers.Logger("error", "In Server: Movie with ID not found")
			msg := helpers.MsgErr(http.StatusNotFound, "not found", "Movie with ID not found")
			helpers.Response(w, http.StatusNotFound, msg)
			return
		}

		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	helpers.Logger("info", "update movies")
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idMovie := params["ID"]

	result, err := handlers.H_DeleteMovie(idMovie)
	if err != nil {

		if gorm.IsRecordNotFoundError(err) {
			helpers.Logger("error", "In Server: Movie with ID not found")
			msg := helpers.MsgErr(http.StatusNotFound, "not found", "Movie with ID not found")
			helpers.Response(w, http.StatusNotFound, msg)
			return
		}

		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	helpers.Logger("info", "deleted movies")
	helpers.Response(w, http.StatusOK, map[string]interface{}{"data": result})
}
