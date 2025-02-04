package vehicle

func NewBike(numberPlate string) *BaseVehicle {
	return &BaseVehicle{
		NumberPlate: numberPlate,
		Type:        BikePark,
	}
}
