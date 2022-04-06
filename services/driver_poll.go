package services

import "fmt"

type service struct {
	drivers []Driver
	booking []Booking
}

func NewDriveService() Services {
	return &service{
		drivers: make([]Driver, 0),
		booking: make([]Booking, 0),
	}
}

func (s *service) RegisterDriver(driverId string) (string, error) {
	//create a new driver and append it to the dirvers list
	//what if the driver is already present? - won't be any duplicates

	newDriver := NewDriver(driverId)
	s.drivers = append(s.drivers, newDriver)

	return fmt.Sprintf("Driver %s registered", driverId), nil
}

//dispactch driver
//find the available driver
//create a booking with the driverId
//append the booking to the services
func (s *service) DispatchDriverForBooking(bookingId string, distance int) (string, error) {

	driver, err := s.FindAvailableDriver()
	if err != nil {
		return "Sorry, driver is not available!", nil
	}

	newBooking := NewBooking(bookingId, driver.GetDriverID(), distance)
	driver.SetDriverAvailability(false)
	s.booking = append(s.booking, newBooking)

	return fmt.Sprintf("Driver %s is assigned to booking %s with %dkm distance", driver.GetDriverID(), newBooking.GetBookingID(), distance), nil
}

	//complete booking
	//if booking id is valid

	//update booking status to end the trip
	//update the driver status to available

//given driver poll
//when completbooking is invoked wiht bookingid
//then i expect invalid booking
	
func (s *service) CompleteBooking(bookingId string) (string, error) {

	booking, err := s.FetchBookingID(bookingId)
	if err != nil {
		return fmt.Sprintf("invalid booking %s", bookingId), nil
	}
	booking.EndTrip() //end the trip
	driverID := booking.GetDriverIDForBooking()

	driver, err := s.FetchDriver(driverID)
	if err != nil {
		return fmt.Sprintf("Invalid driver %s", driverID), nil
	}

	driver.SetDriverAvailability(true) //Chaging the availability to true

	return fmt.Sprintf("Driver %s is released to allocation poll", driverID), nil
}

func (s *service) FindAvailableDriver() (Driver, error) {
	for _, d := range s.drivers {
		if d.CheckAvailability() {
			return d, nil
		}
	}
	return nil, errNoAvailableDriver
}

func (s *service) FetchBookingID(bookingId string) (Booking, error) {
	for _, booking := range s.booking {
		if booking.GetBookingID() == bookingId {
			return booking, nil
		}
	}
	return nil, errInvalidBookingId
}

func (s *service) FetchDriver(driverId string) (Driver, error) {
	for _, driver := range s.drivers {
		if driver.GetDriverID() == driverId {
			return driver, nil
		}
	}
	return nil, errInvalidDriverId
}
