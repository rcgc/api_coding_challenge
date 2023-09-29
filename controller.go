package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// getAll godoc
// @Summary		Get all cars
// @Description Gets all the cars from the database
// @Tags		car
// @Accept		json
// @Produce		json
// @Success		200 		{array} 		Car			"OK"
// @Router		/cars		[get]
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

// getById godoc
// @Summary		Get a car
// @Description	Gets a single car from the database corresponding to the id in the path. Otherwise, returns error
// @Tags		car
// @Accept		json
// @Produce		json
// @Param		id			path			string			true			"Car Id"
// @Success		200			{object}		Car				"OK"
// @Failure		404			{string}		string			"NotFound"
// @Router		/cars/{id} 	[get]
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

// post doc
// @Summary		Create a new car
// @Description	Creates a new car in the database. In case of existing id returns error
// @Tags		car
// @Accept		json
// @Produce		json
// @Param		car			body			Car				true			"Car JSON Object"
// @Success		201			{object}		Car				"OK"
// @Failure		400			{string}		string			"BadRequest"
// @Router		/cars 		[post]
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

// put doc
// @Summary		Update a car
// @Description	Updates an existing car from the database corresponding to the id sent. Otherwise, returns error
// @Tags			car
// @Accept		json
// @Produce		json
// @Param		car			body			Car				true			"Car JSON Object"
// @Success		200			{object}		Car				"OK"
// @Failure		400			{string}		string			"BadRequest"
// @Failure		404			{string}		string
// @Router		/cars		[put]
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

// delete doc
// @Summary		Delete a car
// @Description  Deletes an existing car from the database corresponding to the id in the path. Otherwise, returns error
// @Tags			car
// @Accept		json
// @Produce		json
// @Param		id			path			string			true			"Car Id"
// @Success		204			{string}		string			"NoContent"
// @Failure		404			{string}		string			"NotFound"
// @Router		/cars/{id}	[delete]
func (h *carHandler) delete(w http.ResponseWriter, r *http.Request) {
	defer h.Unlock()
	h.Lock()

	id := idFromUrl(r)

	car := Car{Id: id}

	if id != "-1"{
		q, err := car.deleteCar()

		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		respondWithJSON(w, http.StatusNoContent, q)
		return
	}
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