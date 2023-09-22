package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Car struct {
	Id   string `json:id`
	Make string `json:make`
}

type Cars []Car

type carHandler struct {
	sync.Mutex
	cars Cars
}

func (c *carHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.get(w, r)
	case "POST":
		c.post(w, r)
	case "PUT", "PATCH":
		c.put(w, r)
	case "DELETE":
		c.delete(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "invalid method")
	}
}

func (c *carHandler) get(w http.ResponseWriter, r *http.Request) {
	defer c.Unlock()
	c.Lock()
	id, err := idFromUrl(r)

	if err != nil {
		respondWithJSON(w, http.StatusOK, c.cars)
		return
	}
	if id >= len(c.cars) || id < 0 {
		respondWithError(w, http.StatusNotFound, "not found")
		return
	}
	respondWithJSON(w, http.StatusOK, c.cars[id])
}

func (c *carHandler) post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		respondWithError(w, http.StatusUnsupportedMediaType, "content type 'application/json' required")
		return
	}
	var car Car
	err = json.Unmarshal(body, &car)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer c.Unlock()
	c.Lock()
	c.cars = append(c.cars, car)
	respondWithJSON(w, http.StatusCreated, car)
}

func (c *carHandler) put(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id, err := idFromUrl(r)

	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		respondWithError(w, http.StatusUnsupportedMediaType, "content type 'application/json' required")
		return
	}

	var car Car
	err = json.Unmarshal(body, &car)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer c.Unlock()
	c.Lock()
	if id >= len(c.cars) || id < 0 {
		respondWithError(w, http.StatusNotFound, "not found")
		return
	}

	if car.Make != "" {
		c.cars[id].Make = car.Make
	}

	respondWithJSON(w, http.StatusOK, c.cars[id])
}

func (c *carHandler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := idFromUrl(r)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "not found")
		return
	}
	defer c.Unlock()
	c.Lock()

	if id >= len(c.cars) || id < 0 {
		respondWithError(w, http.StatusNotFound, "not found")
		return
	}
	if id < len(c.cars)-1 {
		c.cars[len(c.cars)-1], c.cars[id] = c.cars[id], c.cars[len(c.cars)-1]
	}
	c.cars = c.cars[:len(c.cars)-1]
	respondWithJSON(w, http.StatusNoContent, "")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func idFromUrl(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.String(), "/")

	if len(parts) != 3 {
		return 0, errors.New("not found")
	}

	id, err := strconv.Atoi(parts[len(parts)-1])

	if err != nil {
		return 0, errors.New("not found")
	}

	return id, nil
}

func newCarHandler() *carHandler {
	return &carHandler{
		cars: Cars{
			Car{"JHk290Xj", "Ford"},
			Car{"fWl37la", "Toyota"},
			Car{"1i3xjRllc", "Toyota"},
			Car{"dku43920s", "Ford"},
		},
	}
}