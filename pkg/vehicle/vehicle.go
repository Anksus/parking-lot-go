package vehicle

type VehicleType int

const (
	CarPark VehicleType = iota
	BikePark
	TruckPark
)

type BaseVehicle struct {
	NumberPlate string
	Type        VehicleType
}

func (v *BaseVehicle) GetVehicleType() VehicleType {
	return v.Type
}
