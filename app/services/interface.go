package services

import "fmt"

type Vehicle interface {
	Start() string
	Stop() string
}

type Car struct {
	Model string
}

func (c Car) Start() string {
	return fmt.Sprintf("Car %s is moving...", c.Model)
}

func (c Car) Stop() string {
	return fmt.Sprintf("Car %s is stopping...", c.Model)
}

type Bike struct {
	Brand string
}

func (b Bike) Start() string {
	return fmt.Sprintf("Bike %s is moving...", b.Brand)
}

func (b Bike) Stop() string {
	return fmt.Sprintf("Bike %s is stopping...", b.Brand)
}

func OperateVehicle(v Vehicle) {
	fmt.Println(v.Start())
	fmt.Println(v.Stop())
}
