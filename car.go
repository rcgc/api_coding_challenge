package main

// swagger:model Car
type Car struct {
	// ID of car
	// in: string 
	Id       string  `json:id`
	// Make of car
	// in: string
	Make     string  `json:make`
	// Model of car
	// in: string
	Model    string  `json:model`
	// Package of car
	// in: string
	Package  string  `json:package`
	// Color of car
	// in: string
	Color    string  `json:color`
	// Year of car
	// in: int
	Year     int     `json:year`
	// Category of car
	// in: string
	Category string  `json:category`
	// Mileage of car
	// in: float64
	Mileage  float64 `json:mileage`
	// Price of car
	// in: float64
	Price    float64 `json:price`
}

var db Db

var m carMiddleware

func (c *Car) getAllCars() ([]Car, error) {
	cars, err := db.getAll()
	if err != nil {
		return []Car{}, err
	}
	return cars, err
}

func (c *Car) getCarById() (Car, error) {
	err := m.validate_getById(c)
	if err != nil {
		return Car{}, err
	}

	car, err := db.getById(c.Id)

	if err != nil {
		return Car{}, err
	}

	return car, nil
}

func (c *Car) createCar() (Car, error) {
	err := m.validate_create(c)
	if err != nil {
		return Car{}, err
	}

	car, err := db.add(c)

	if err != nil {
		return Car{}, err
	}

	return car, nil
}

func (c *Car) updateCar() (Car, error) {
	err := m.validate_update(c)
	if err != nil {
		return Car{}, err
	}

	car, err := db.update(c)

	if err != nil {
		return Car{}, err
	}

	return car, nil
}

func (c *Car) deleteCar() (Car, error) {
	err := m.validate_delete(c)
	if err != nil {
		return Car{}, err
	}

	car, err := db.delete(c.Id)

	if err != nil {
		return Car{}, err
	}

	return car, nil
}