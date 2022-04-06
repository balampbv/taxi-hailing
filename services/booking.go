package services

type booking struct {
	id           string
	driverId     string
	tripDistance int
	ongoing      bool
}

type Booking interface {
	EndTrip() error
	GetBookingID() string
	GetDriverIDForBooking() string
	//
}

func NewBooking(bookingID string, driverID string, distance int) Booking {
	return &booking{
		id:           bookingID, //Todo auto-increment
		driverId:     driverID,
		tripDistance: distance,
		ongoing:      true,
	}
}

func (b *booking) EndTrip() error {
	if b.ongoing {
		b.ongoing = false
		return nil
	}
	return errorTripEnded
}

func (b *booking) GetBookingID() string {
	return b.id
}

func (b *booking) GetDriverIDForBooking() string {
	return b.driverId
}
