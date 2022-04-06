package services

import "errors"

const (
	errTripHasEnded   = "Trip already ended"
	invalidBookingId  = "Invalid Booking ID"
	invalidDriverId   = "Invalid Driver ID"
	noAvailableDriver = "No available driver found"
)

var (
	errorTripEnded       = errors.New(errTripHasEnded)
	errNoAvailableDriver = errors.New(noAvailableDriver)
	errInvalidBookingId  = errors.New(invalidBookingId)
	errInvalidDriverId   = errors.New(invalidDriverId)
)
