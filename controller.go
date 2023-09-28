package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// getAll godoc
// @Summary List cars
// @Description Retrieves all the cars stored in the database
// @Tags car
// @Accept  json
// @Produce  json
// @Success 200 {array} []Car
// @Router /cars [get]
func (h *carHandler) getAll(w http.ResponseWriter, r *http.Request){
	defer h.Unlock()
	h.Lock()

	car := Car{}
	q, err := car.getAllCars()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	respondWithJSON(w, http.StatusOK, q)
}

func (h *carHandler) getById(w http.ResponseWriter, r *http.Request) {
	defer h.Unlock()
	h.Lock()


	id := idFromUrl(r)

	car := Car{Id: id}
	if id != "-1" {
		query, err := car.getCarById()
		
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, query)
		return
	}
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

		if err != nil{
			if err.Error() == "id field empty" || err.Error() == "make field empty" || err.Error() == "model field empty" || 
			err.Error() == "package field empty" || err.Error() == "color field empty" || err.Error() == "year field must be gt 0" ||
			err.Error() == "category field empty" || err.Error() == "mileage field must be gt 0" || err.Error() == "price field must be gt 0" {
				respondWithError(w, http.StatusBadRequest, err.Error())
				return
			}
			if err.Error() == "id not found" {
				respondWithError(w, http.StatusNotFound, err.Error())
				return
			}
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
	w.Header().Add("authorization", "Access-Control-Allow-Origin")
	w.WriteHeader(code)
	w.Write(response)
}

func idFromUrl(r *http.Request) (string) {
	parts := strings.Split(r.URL.String(), "/")

	if len(parts) < 3 {
		return "-1"
	}

	if parts[2] == ""{
		return "-1"
	}
	
	/*
	id, err := strconv.Atoi(parts[len(parts)-1])

	if err != nil {
		return -1, errors.New("not found")
	}
	*/

	id :=  parts[2]

	return id
}

/*
func splitUrlParameters(r *http.Request) ([]string){
	parts := strings.Split(r.URL.String(), "/")
	return parts
}
*/