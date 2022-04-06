package main

import (
	"bufio"
	"fmt"
	"gojek-drive/services"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fetch inputs from cmd
	//process each commands
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------")

	driveService := services.NewDriveService()
	count := 0
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		fmt.Println(text)
		inputArr := strings.Split(text, " ")
		switch inputArr[0] {
		case "register_driver":
			driverId := inputArr[1]
			result, _ := driveService.RegisterDriver(driverId)
			fmt.Println(result)
		case "dispatch_driver_for_a_booking":
			distanceStr := inputArr[1]
			distance, _ := strconv.Atoi(distanceStr)
			count++
			result, _ := driveService.DispatchDriverForBooking("B"+strconv.Itoa(count), distance)
			fmt.Println(result)
		case "complete_booking":
			bookingId := inputArr[1]
			result, _ := driveService.CompleteBooking(bookingId)
			fmt.Println(result)

		}

	}
}

//register_driver <<driver_id>>
//$ dispatch_driver_for_a_booking <<booking distance>>
//$ complete_booking <<booking-id>>

//DriverPoll - max //first available dirver
//Driver  - id, availabilty
//Booking - id, distance, driverId
//Register driver
//create booking or dispatch
//complete booking
