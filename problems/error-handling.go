package main

import (
	"errors"
	"fmt"
)

type Truck struct {
	id string
}

func (t *Truck) LoadCargo() error {
	return errors.New("not implemented")
}

func processTruck(truck Truck) {
	fmt.Println("Processing truck", truck.id)

	var err error
	if truck.id == "Truck-1" {
		err = truck.LoadCargo()
	}
	if err != nil {
		fmt.Println("Error loading cargo: ", err)
		return
	}

	fmt.Println("Cargo loaded successfully")
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Println("Truck arrived: ", truck.id)
		processTruck(truck)
	}
}
