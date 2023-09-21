package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-xsis-movie/controllers"
	"go-xsis-movie/models/entities"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var (
	server *httptest.Server
)

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/Movies", controllers.C_AddMovie).Methods("POST")
	r.HandleFunc("/Movies", controllers.C_GetAllMovie).Methods("GET")
	r.HandleFunc("/Movies/{ID}", controllers.C_GetSingleMovieId).Methods("GET")
	r.HandleFunc("/Movies/{ID}", controllers.C_UpdateSingleMovieId).Methods("PATCH")
	r.HandleFunc("/Movies/{ID}", controllers.C_DeleteMovie).Methods("DELETE")
	return r
}

func TestMain(m *testing.M) {
	server = httptest.NewServer(Router())
	defer server.Close()
	code := m.Run()
	os.Exit(code)
}

func Test_AddMovie(t *testing.T) {
	payload := entities.Movie{
		Title:       "Pengabdi Setan 2 Comunion",
		Description: "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
		Rating:      9,
		Image:       "bca",
	}
	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "http://localhost:7000/Movies", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status Created; got: %v", res.Status)
	}

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	fmt.Printf("Response from Test_AddMovie:\n%v\n", response)
}

func Test_ReadAllMovie(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:7000/Movies", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got: %v", res.Status)
	}

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	fmt.Printf("Response from Test_ReadAllMovie:\n%v\n", response)
}

func Test_ReadSingleMovieId(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:7000/Movies/1", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got: %v", res.Status)
	}

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	fmt.Printf("Response from Test_ReadSingleMovieId:\n%v\n", response)
}

func Test_UpdateSingleMovieId(t *testing.T) {
	updateData := map[string]interface{}{
		"title":       "Pengabdi setan 2 Comunion",
		"description": "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
		"rating":      12,
		"image":       "xyz",
	}

	updateDataJSON, err := json.Marshal(updateData)
	if err != nil {
		t.Fatalf("could not marshal update data: %v", err)
	}

	server := httptest.NewServer(Router())
	defer server.Close()

	req, err := http.NewRequest("PATCH", "http://localhost:7000/Movies/1", bytes.NewReader(updateDataJSON))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got: %v", res.Status)
	}

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	fmt.Printf("Response from Test_UpdateSingleMovieId:\n%v\n", response)
}

func Test_DeleteMovie(t *testing.T) {

	req, err := http.NewRequest("DELETE", "http://localhost:7000/Movies/1", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status Not Found; got: %v", res.Status)
	}

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	fmt.Printf("Response from Test_DeleteMovie:\n%v\n", response)
}
