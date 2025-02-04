package vehicle

func NewCar(numberPlate string) *BaseVehicle {
	return &BaseVehicle{
		NumberPlate: numberPlate,
		Type:        CarPark,
	}
}
