package services

type driver struct {
	id          string
	availabilty bool
	tripCount   int
}

//Dirver interface exposes the methods related to the driver object
type Driver interface {
	GetDriverID() string
	CheckAvailability() bool
	SetDriverAvailability(bool)
	//IncrementTrip
}

func NewDriver(driverId string) Driver {
	return &driver{
		id:          driverId,
		availabilty: true,
	}
}

func (d *driver) CheckAvailability() bool {
	return d.availabilty
}

func (d *driver) GetDriverID() string {
	return d.id
}

func (d *driver) SetDriverAvailability(availability bool) {
	d.availabilty = availability
}
