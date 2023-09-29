package main

import "fmt"

type Db struct {
	cars []Car
}

func (db *Db) getAll() ([]Car, error) {
	return db.cars, nil
}

func (db *Db) getById(id string) (Car, error) {

	for _, v := range db.cars {
		if v.Id == id {
			return v, nil
		}
	}

	return Car{}, fmt.Errorf("id not found")
}

func (db *Db) add(c *Car) (Car, error){
	car := Car{Id: c.Id, Make: c.Make, Model: c.Model, Package: c.Package, Color: c.Color, Year: c.Year, Category: c.Category, Mileage: c.Mileage, Price: c.Price}

	for _, v := range db.cars{
		if car.Id == v.Id {
			return car, fmt.Errorf("id already exists")
		}
	}

	db.cars = append(db.cars, car)
	return car, nil
}

func (db *Db) update(c *Car) (Car, error) {
	car := Car{Id: c.Id, Make: c.Make, Model: c.Model, Package: c.Package, Color: c.Color, Year: c.Year, Category: c.Category, Mileage: c.Mileage, Price: c.Price}
	
	for i, v := range db.cars{
		if v.Id == car.Id {
			db.cars[i].Make = car.Make
			db.cars[i].Model = car.Model
			db.cars[i].Package = car.Package
			db.cars[i].Color = car.Color
			db.cars[i].Year = car.Year
			db.cars[i].Category = car.Category
			db.cars[i].Mileage = car.Mileage
			db.cars[i].Price = car.Price
			return car, nil 
		}
	}
	return Car{}, fmt.Errorf("id not found")
}

func (db *Db) delete(id string) (Car, error){
	index := -1

	for i, v := range db.cars {
		if v.Id == id {
			index = i
		}
	}

	if index != -1 {
		if index < len(db.cars)-1 {
			db.cars[len(db.cars)-1], db.cars[index] = db.cars[index], db.cars[len(db.cars)-1]
		}
		db.cars = db.cars[:len(db.cars)-1]
		return Car{}, nil
	}

	return Car{}, fmt.Errorf("id not found")
}