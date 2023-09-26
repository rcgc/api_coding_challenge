package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

type carHandler struct {
	sync.Mutex
}

func (h *carHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
	case "POST":
		h.post(w, r)
	case "PUT", "PATCH":
		h.put(w, r)
	case "DELETE":
		h.delete(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "invalid method")
	}
}

func (h *carHandler) get(w http.ResponseWriter, r *http.Request) {
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
	if r.URL.String() == "/cars/profile" || r.URL.String() == "/cars/profile/" {
		err = json.Unmarshal(body, &car)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	defer h.Unlock()
	h.Lock()

	if r.URL.String() == "/cars" || r.URL.String() == "/cars/" {
		q, err := car.getAllCars()

		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		
		respondWithJSON(w, http.StatusOK, q)
		return
	} else if r.URL.String() == "/cars/profile" || r.URL.String() == "/cars/profile/" {
		query, err := car.getCarById()
		
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, query)
		return
	}
	respondWithError(w, http.StatusBadRequest, "no valid URL")
}

func (h *carHandler) post(w http.ResponseWriter, r *http.Request) {
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
	if r.URL.String() == "/cars" || r.URL.String() == "/cars/"{
		err = json.Unmarshal(body, &car)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer h.Unlock()
		h.Lock()
		q, err := car.createCar()

		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondWithJSON(w, http.StatusCreated, q)
		return
	}
	respondWithError(w, http.StatusBadRequest, "no valid URL")
}

func (h *carHandler) put(w http.ResponseWriter, r *http.Request) {
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
	if r.URL.String() == "/cars" || r.URL.String() == "/cars/"{
		var car Car
		err = json.Unmarshal(body, &car)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		defer h.Unlock()
		h.Lock()

		q, err := car.updateCar()

		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, q)
		return
	}
	respondWithError(w, http.StatusBadRequest, "no valid URL")
}

func (h *carHandler) delete(w http.ResponseWriter, r *http.Request) {
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
	if r.URL.String() == "/cars" || r.URL.String() == "/cars/"{
		var car Car
		err = json.Unmarshal(body, &car)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		defer h.Unlock()
		h.Lock()

		q, err := car.deleteCar()

		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		respondWithJSON(w, http.StatusNoContent, q)
		return
	}
	respondWithError(w, http.StatusBadRequest, "no valid URL")
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

/*
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
*/

/*
func splitUrlParameters(r *http.Request) ([]string){
	parts := strings.Split(r.URL.String(), "/")
	return parts
}
*/

func newCarHandler() *carHandler {
	
	preload := []Car{
		{"JHk290Xj",	"Ford",		"F10",		"Base",		"Silver",		2010,	"Truck",	120123,		1999900}, 
		{"fWl37la",		"Toyota",	"Camry",	"SE",		"White",		2019,	"Sedan",	3999,		2899000},
		{"1i3xjRllc",	"Toyota",	"Rav4",		"XSE",		"Red",			2018,	"SUV",		24001,		2275000},
		{"dku43920s",	"Ford",		"Bronco",	"Badlands",	"Burnt Orange",	2022,	"SUV",		1,			4499000},
	}

	for _, v := range preload {
		v.createCar()
	}
	
	return &carHandler{}
}