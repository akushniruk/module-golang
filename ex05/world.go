package barber

import "fmt"

func world ( fromShopChannel <-chan *Customer , terminationChannel chan<- *Customer ) {
	customersTurnedAway := 0
	customersTrimmed := 0

	for exists := true ; exists ; {
		customer := <- fromShopChannel

		if customer.id == -1 {
			exists = false
			fmt.Printf ( "\nTrimmed %d and turned away %d today.\n" ,  customersTrimmed , customersTurnedAway )
			terminationChannel <- customer
		} else {
			if customer.done {
				customersTrimmed++
			} else {
				customersTurnedAway++
			}
			fmt.Printf ( "Customer %d left the shop.\n" , customer.id )
		}
	}
}