package vehicle

func NewTruck(numberPlate string) *BaseVehicle {
	return &BaseVehicle{
		NumberPlate: numberPlate,
		Type:        TruckPark,
	}
}
