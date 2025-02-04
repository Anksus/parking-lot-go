package main

import (
	"Anksus/parking-lot-go/pkg/parkinglot"
	"Anksus/parking-lot-go/pkg/vehicle"
	"fmt"
)

func main() {
	pl := parkinglot.NewConcreteParkingLotBuilder(3).AddCarPark(2).AddTruckPark(2).
		AddBikePark(2).Build()
	pl.DisplayAvailability()

	car1 := vehicle.NewCar("KA-01-HH-1234")
	car2 := vehicle.NewCar("KA-01-HH-9999")

	//bike1 := vehicle.NewBike("KA-01-HH-1234")
	//
	//truck1 := vehicle.NewTruck("KA-01-HH-1234")

	err := pl.ParkVehicle(car1)
	err = pl.ParkVehicle(car2)
	if err != nil {
		fmt.Println(err)
	}
	pl.DisplayAvailability()

	err = pl.UnPark(car1)
	if err != nil {
		fmt.Println(err)
	}
	pl.DisplayAvailability()
}
