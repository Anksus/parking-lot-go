package parkinglot

import (
	"Anksus/parking-lot-go/pkg/level"
	"Anksus/parking-lot-go/pkg/vehicle"
	"errors"
	"fmt"
)

type ParkingLot struct {
	Levels map[int]*level.Level
}

func (pl *ParkingLot) ParkVehicle(v *vehicle.BaseVehicle) error {
	for k := range pl.Levels {
		if isParked := pl.Levels[k].ParkVehicle(v); isParked {
			return nil
		}
	}

	return errors.New(fmt.Sprintf("vehicle: %+v, can not be parked", v))
}

func (pl *ParkingLot) DisplayAvailability() {
	carPark := 0
	bikePark := 0
	truckPark := 0

	for _, v := range pl.Levels {
		carPark += v.CarPark
		bikePark += v.BikePark
		truckPark += v.TruckPark
	}
	fmt.Println(fmt.Sprintf("Available car parks: %d, bike parks: %d, truct parks: %d", carPark, bikePark, truckPark))
}

func (pl *ParkingLot) UnPark(v *vehicle.BaseVehicle) error {
	for k := range pl.Levels {
		if unParked := pl.Levels[k].UnPark(v); unParked {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("can't exit this vehicle, %+v", v))
}
