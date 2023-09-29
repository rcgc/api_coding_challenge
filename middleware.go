package main

import "fmt"

type carMiddleware struct {}

func (m *carMiddleware) validate_getById(c *Car) error {
	if c.Id == "" {
		return fmt.Errorf("id field empty")
	}

	return nil
}

func (m *carMiddleware)validate_create(c *Car) error {
	if c.Id == "" {
		return fmt.Errorf("id field empty")
	}
	if c.Make == "" {
		return fmt.Errorf("make field empty")
	}
	if c.Model == "" {
		return fmt.Errorf("model field empty")
	}
	if c.Package == "" {
		return fmt.Errorf("package field empty")
	}
	if c.Color == "" {
		return fmt.Errorf("color field empty")
	}
	if c.Year <= 0 {
		return fmt.Errorf("year field must be gt 0")
	}
	if c.Category == "" {
		return fmt.Errorf("category field empty")
	}
	if c.Mileage < 0 {
		return fmt.Errorf("mileage field must be ge 0")
	}
	if c.Price <= 0 {
		return fmt.Errorf("price field must be gt 0")
	}

	return nil
}

func (m *carMiddleware) validate_update(c *Car) error {
	if c.Id == "" {
		return fmt.Errorf("id field empty")
	}
	if c.Make == "" {
		return fmt.Errorf("make field empty")
	}
	if c.Model == "" {
		return fmt.Errorf("model field empty")
	}
	if c.Package == "" {
		return fmt.Errorf("package field empty")
	}
	if c.Color == "" {
		return fmt.Errorf("color field empty")
	}
	if c.Year <= 0 {
		return fmt.Errorf("year field must be gt 0")
	}
	if c.Category == "" {
		return fmt.Errorf("category field empty")
	}
	if c.Mileage < 0 {
		return fmt.Errorf("mileage field must be ge 0")
	}
	if c.Price <= 0 {
		return fmt.Errorf("price field must be gt 0")
	}

	return nil
}

func (m *carMiddleware) validate_delete(c *Car) error {
	if c.Id == "" {
		return fmt.Errorf("id field empty")
	}

	return nil
}