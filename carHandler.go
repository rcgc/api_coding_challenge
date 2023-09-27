package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type carHandler struct {
	sync.Mutex
}

func (h *carHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if idFromUrl(r) == "-1"{
			h.getAll(w, r)
		} else {
			h.getById(w, r)
		}
		
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

// swagger:route GET /cars Car getAllCars
//
// Gets all the cars from the database
//
// ---
// Consumes:
// - application/json 
//
// Produces:
// - application/json
//
// Schemes: http
//
// Responses:
// 200: []Car
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

// swagger:route GET /cars/{id} Car getCarById
// 
// Gets a car by id from the database
//
// ---
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Schemes: http
//
// Parameters:
// + id: id
//   in: path
//   description: id of the car
//   required: true
//   type: string
//
// Responses:
// 200: Car
// 404:
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

// swagger:route POST /cars Car createCar
// 
// Creates a new Car in the database, in case of similar Id returns error
//
// ---
// Responses:
// 201: Car
// 400:
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

// swagger:route PUT /cars Car updateCar
// 
// Updates an existing Car in the database according to the Id sent, otherwise returns error
// 
// ---
// Responses:
// 200: Car
// 400:
// 404:
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

// swagger:route DELETE /cars Car deleteCar
//
// Deletes an existing Car in the database according to the Id sent, otherwise returns error
//
// ---
// Responses:
// 204:
// 404:
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