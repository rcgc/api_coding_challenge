package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
)

var (
	listCars  = regexp.MustCompile(`^\/cars[\/]*$`)
	getCar    = regexp.MustCompile(`^\/cars\/(\d+)$`)
)

type car struct {
	Id			string		`json:"id"`
	Make		string		`json:"make"`
	Model		string		`json:"model"`
	Package		string		`json:"package"`
	Color		string		`json:"color"`
	Year		int			`json:"year"`
	Category	string		`json:"category"`
	Mileage		int			`json:"mileage"`
	Price		float64		`json:"price"`
}

type datastore struct{
	m map[string]car
	*sync.RWMutex
}

type carHandler struct {
	store *datastore
}

func (h *carHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	switch{
	case r.Method == http.MethodGet && listCars.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getCar.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	default:
		notFound(w, r)
		return
	}	
}

func (h *carHandler) List(w http.ResponseWriter, r *http.Request){
	h.store.RLock()
	cars := make([]car, 0, len(h.store.m))

	for _, v:= range h.store.m {
		cars = append(cars, v)
	}

	h.store.RUnlock()
	jsonBytes, err := json.Marshal(cars)
	
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *carHandler) Get(w http.ResponseWriter, r *http.Request){
	matches := getCar.FindStringSubmatch(r.URL.Path)
	
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	c, ok := h.store.m[matches[1]]
	h.store.RUnlock()

	if !ok{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("car not found"))
		return
	}

	jsonBytes, err := json.Marshal(c)
	if err != nil {
		internalServerError(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func internalServerError(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func main() {
	mux := http.NewServeMux()
	carH := &carHandler{
		store: &datastore{
			m: map[string]car{
				"1": {Id: "JHK290Xj", 	Make: "Ford", 	Model: "F10", 		Package: "Base", 		Color: "Silver", 		Year: 2010, 	Category: "Truck", 	Mileage: 120123, 	Price: 1999900},
				"2": {Id: "fWl37la", 	Make: "Toyota", Model: "Camry", 	Package: "SE", 			Color: "White", 		Year: 2019,		Category: "Sedan", 	Mileage: 3999, 		Price: 2899000},
				"3": {Id: "1i3xjRllc", 	Make: "Toyota", Model: "Rav4", 		Package: "XSE", 		Color: "Red", 			Year: 2018, 	Category: "SUV", 	Mileage: 24001, 	Price: 2275000},
				"4": {Id: "dku43920s", 	Make: "Ford",	Model: "Bronco",	Package: "Badlands", 	Color: "Burnt Orange", 	Year: 2022, 	Category: "SUV", 	Mileage: 1, 		Price: 4499000},
			},
			RWMutex: &sync.RWMutex{},
		},
	}

	mux.Handle("/cars", carH)
	mux.Handle("/cars/", carH)

	http.ListenAndServe(":8080", mux)
}