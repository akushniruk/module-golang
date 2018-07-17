package barber

import (
	"fmt"
	"math/rand"
	"time"
)

type Customer struct {
	id int ;
	done bool
}

func barber ( hairTrimTime func ( ) time.Duration , fromShopChannel <-chan *Customer , toShopChannel chan<- *Customer ) {
	customersTrimmed := 0

	for working := true ; working ; {
		customer := <- fromShopChannel

		if customer.id == -1 {
			working = false
			fmt.Printf ( "Knocking off time, trimed %d customers.\n" , customersTrimmed )
		} else {
			fmt.Printf ( "Starting Customer %d\n" , customer.id )
			time.Sleep ( hairTrimTime ( ) )
			customersTrimmed ++
			customer.done = true
			fmt.Printf ( "Finished Customer %d\n" , customer.id )
		}

		toShopChannel <- customer
	}
}

func shop ( numberOfSeats int , fromWorldChannel <-chan *Customer , toBarberChannel chan<- *Customer ,
	fromBarberChannel <-chan *Customer , toWorldChannel chan<- *Customer ) {

	seatsFilled := 0
	customersTurnedAway := 0
	customersTrimmed := 0

	for isOpen := true ; isOpen ; {
		var customer *Customer

		select {
		case customer = <- fromWorldChannel :
			if customer.id == -1 {
				toBarberChannel <- customer
			} else {
				if seatsFilled <= numberOfSeats {
					seatsFilled++
					fmt.Printf ( "Customer %d takes a seat. %d in use.\n" , customer.id , seatsFilled )
					toBarberChannel <- customer
				} else {
					customersTurnedAway++
					fmt.Printf ( "Customer %d turned away.\n" , customer.id )
					toWorldChannel <- customer
				}
			}
		case customer = <- fromBarberChannel :
			if customer.id == -1 {
				isOpen = false
				fmt.Printf ( "Shop closing, %d trimmed and %d turned away.\n" ,  customersTrimmed , customersTurnedAway )
			} else {
				customersTrimmed++
				seatsFilled--
				fmt.Printf ( "Customer %d leaving trimmed.\n" , customer.id )
			}

			toWorldChannel <- customer
		}
	}
}
