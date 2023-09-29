package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -coverprofile="coverage.out"
// go tool cover -html="coverage.out"

func TestCreateCar_WhenSuccessful(t *testing.T) {
	car := Car{ Id: "abcdefghi", Make: "Nissan", Model: "March", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.createCar()

	assert.Equal(t, err, nil)
}

func TestCreateCar_WhenIdAlreadyExistsInDb(t *testing.T){
	car := Car{ Id: "rstuvwxyz", Make: "Nissan", Model: "Altima", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	car.createCar()
	_, err := car.createCar()

	assert.Equal(t, err.Error(), "id already exists")
}

func TestCreateCar_WhenIdFieldEmpty(t *testing.T){
	car := Car{ Id: "", Make: "Nissan", Model: "Altima", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.createCar()

	assert.Equal(t, err.Error(), "id field empty")
}

func TestCreateCar_WhenMakeFieldEmpty(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "", Model: "Altima", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.createCar()

	assert.Equal(t, err.Error(), "make field empty")
}

func TestCreateCar_WhenModelFieldEmpty(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.createCar()

	assert.Equal(t, err.Error(), "model field empty")
}

func TestCreateCar_WhenPackageFieldEmpty(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "Altima", Package: "", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.createCar()

	assert.Equal(t, err.Error(), "package field empty")
}

func TestCreateCar_WhenColorFieldEmpty(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "Altima", Package: "XX", Color: "", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.createCar()

	assert.Equal(t, err.Error(), "color field empty")
}

func TestCreateCar_WhenYearFieldLE_0(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "Altima", Package: "XX", Color: "Gray", Year: 0, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.createCar()

	assert.Equal(t, err.Error(), "year field must be gt 0")
}

func TestCreateCar_WhenCategoryFieldEmpty(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "Altima", Package: "XX", Color: "Gray", Year: 2013, Category: "", Mileage: 799, Price: 2499000 }
	_, err := car.createCar()

	assert.Equal(t, err.Error(), "category field empty")
}

func TestCreateCar_WhenMileageFieldLE_0(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "Altima", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: -1, Price: 2499000 }
	_, err := car.createCar()

	assert.Equal(t, err.Error(), "mileage field must be ge 0")
}

func TestCreateCar_WhenPriceFieldLE_0(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "Altima", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 0 }
	_, err := car.createCar()

	assert.Equal(t, err.Error(), "price field must be gt 0")
}

func TestGetAllCars_WhenSuccessful(t *testing.T){
	var car Car
	_, err := car.getAllCars()

	assert.Equal(t, err, nil)
}

func TestGetCarById_WhenSuccessful(t *testing.T){
	car := Car{ Id: "jklmnopqr", Make: "Nissan", Model: "Ultra", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	car.createCar()
	_, err := car.getCarById()

	assert.Equal(t, err, nil)
}

func TestGetCarById_WhenIdNoFound(t *testing.T){
	car := Car{ Id: "zzzzzzzzz"}
	_, err := car.getCarById()

	assert.Equal(t, err.Error(), "id not found")
}

func TestGetCarById_WhenIdFieldEmpty(t *testing.T){
	car := Car{}
	_, err := car.getCarById()

	assert.Equal(t, err.Error(), "id field empty")
}

func TestUpdateCar_WhenSuccessful(t *testing.T){
	car := Car{ Id: "abcdefghi", Make: "Nissan", Model: "Versa", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.updateCar()

	assert.Equal(t, err, nil)
}

func TestUpdateCar_WhenIdNotFound(t *testing.T){
	car := Car{ Id: "zzzzzzzzz", Make: "Nissan", Model: "Versa", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.updateCar()

	assert.Equal(t, err.Error(), "id not found")
}

func TestUpdateCar_WhenIdFieldEmpty(t *testing.T){
	car := Car{ Id: "", Make: "Nissan", Model: "Altima", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.updateCar()

	assert.Equal(t, err.Error(), "id field empty")
}

func TestUpdateCar_WhenMakeFieldEmpty(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "", Model: "Altima", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.updateCar()

	assert.Equal(t, err.Error(), "make field empty")
}

func TestUpdateCar_WhenModelFieldEmpty(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.updateCar()

	assert.Equal(t, err.Error(), "model field empty")
}

func TestUpdateCar_WhenPackageFieldEmpty(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "Altima", Package: "", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.updateCar()

	assert.Equal(t, err.Error(), "package field empty")
}

func TestUpdateCar_WhenColorFieldEmpty(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "Altima", Package: "XX", Color: "", Year: 2013, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.updateCar()

	assert.Equal(t, err.Error(), "color field empty")
}

func TestUpdateCar_WhenYearFieldLE_0(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "Altima", Package: "XX", Color: "Gray", Year: 0, Category: "SUV", Mileage: 799, Price: 2499000 }
	_, err := car.updateCar()

	assert.Equal(t, err.Error(), "year field must be gt 0")
}

func TestUpdateCar_WhenCategoryFieldEmpty(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "Altima", Package: "XX", Color: "Gray", Year: 2013, Category: "", Mileage: 799, Price: 2499000 }
	_, err := car.updateCar()

	assert.Equal(t, err.Error(), "category field empty")
}

func TestUpdateCar_WhenMileageFieldLE_0(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "Altima", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: -1, Price: 2499000 }
	_, err := car.updateCar()

	assert.Equal(t, err.Error(), "mileage field must be ge 0")
}

func TestUpdateCar_WhenPriceFieldLE_0(t *testing.T){
	car := Car{ Id: "opqrstuvw", Make: "Nissan", Model: "Altima", Package: "XX", Color: "Gray", Year: 2013, Category: "SUV", Mileage: 799, Price: 0 }
	_, err := car.updateCar()

	assert.Equal(t, err.Error(), "price field must be gt 0")
}

func TestDeleteCar_WhenSuccessful(t *testing.T){
	car := Car{ Id: "abcdefghi" }
	_, err := car.deleteCar()

	assert.Equal(t, err, nil)
}

func TestDeleteCar_WhenIdNotFound(t *testing.T){
	car := Car{ Id: "zzzzzzzzz" }
	_, err := car.deleteCar()

	assert.Equal(t, err.Error(), "id not found")
}

func TestDeleteCar_WhenIdFieldEmpty(t *testing.T){
	car := Car{ Id:"" }
	_, err := car.deleteCar()

	assert.Equal(t, err.Error(), "id field empty")
}