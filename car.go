package main

// Car model info
// @Description car information
type Car struct {
	Id       string  `json:id`
	Make     string  `json:make`
	Model    string  `json:model`
	Package  string  `json:package`
	Color    string  `json:color`
	Year     int     `json:year`
	Category string  `json:category`
	Mileage  float64 `json:mileage`
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

	_, err = db.delete(c.Id)

	if err != nil {
		return Car{}, err
	}

	return Car{}, nil
}