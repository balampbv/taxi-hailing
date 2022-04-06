package services

type Services interface {
	RegisterDriver(driverId string) (string, error)
	DispatchDriverForBooking(bookingId string, distance int) (string, error)
	CompleteBooking(bookingID string) (string, error)
}

//register_driver <<driver_id>>
//$ dispatch_driver_for_a_booking <<booking distance>>
//$ complete_booking <<booking-id>>
