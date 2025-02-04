package level

import "Anksus/parking-lot-go/pkg/vehicle"

type Level struct {
	CarPark   int
	BikePark  int
	TruckPark int
	Vehicles  []*vehicle.BaseVehicle
}

func (l *Level) ParkVehicle(v *vehicle.BaseVehicle) bool {
	if v.GetVehicleType() == vehicle.CarPark && l.CarPark > 0 {
		l.Vehicles = append(l.Vehicles, v)
		l.CarPark--
		return true
	} else if v.GetVehicleType() == vehicle.BikePark && l.BikePark > 0 {
		l.Vehicles = append(l.Vehicles, v)
		l.BikePark--
		return true
	} else if v.GetVehicleType() == vehicle.TruckPark && l.TruckPark > 0 {
		l.Vehicles = append(l.Vehicles, v)
		l.TruckPark--
		return true
	}
	return false
}

func (l *Level) UnPark(v *vehicle.BaseVehicle) bool {
	for i, veh := range l.Vehicles {
		if veh.NumberPlate == v.NumberPlate {
			l.Vehicles = append(l.Vehicles[:i], l.Vehicles[i+1:]...)
			if v.GetVehicleType() == vehicle.CarPark {
				l.CarPark++
			} else if v.GetVehicleType() == vehicle.BikePark {
				l.BikePark++
			} else if v.GetVehicleType() == vehicle.TruckPark {
				l.TruckPark++
			}
			return true
		}
	}
	return false
}
