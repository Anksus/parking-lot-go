package parkinglot

import (
	"Anksus/parking-lot-go/pkg/level"
)

type ParkingLotBuilder interface {
	AddCarPark(carParkNum int) ParkingLotBuilder
	AddBikePark(bikeParkNum int) ParkingLotBuilder
	AddTruckPark(truckParkNum int) ParkingLotBuilder
	Build() *ParkingLot
}

type ConcreteParkingLotBuilder struct {
	ParkingLot *ParkingLot
}

func NewConcreteParkingLotBuilder(levelNum int) *ConcreteParkingLotBuilder {
	mp := map[int]*level.Level{}
	for i := 0; i < levelNum; i++ {
		mp[i] = &level.Level{}
	}
	return &ConcreteParkingLotBuilder{
		ParkingLot: &ParkingLot{Levels: mp},
	}
}

func (cpl *ConcreteParkingLotBuilder) AddCarPark(carParkNum int) ParkingLotBuilder {
	for k := range cpl.ParkingLot.Levels {
		cpl.ParkingLot.Levels[k].CarPark = carParkNum
	}
	return cpl
}

func (cpl *ConcreteParkingLotBuilder) AddBikePark(bikeParkNum int) ParkingLotBuilder {
	for k := range cpl.ParkingLot.Levels {
		cpl.ParkingLot.Levels[k].BikePark = bikeParkNum
	}
	return cpl
}

func (cpl *ConcreteParkingLotBuilder) AddTruckPark(truckParkNum int) ParkingLotBuilder {
	for k := range cpl.ParkingLot.Levels {
		cpl.ParkingLot.Levels[k].TruckPark = truckParkNum
	}
	return cpl
}

func (cpl *ConcreteParkingLotBuilder) Build() *ParkingLot {
	return cpl.ParkingLot
}
