package barber

import (
	"fmt"
	"time"
)

func runSimulation ( numberOfCustomers , numberOfSeats int , nextCustomerWaitTime , hairTrimTime func ( ) time.Duration ) {
	worldToShopChannel := make ( chan *Customer )
	shopToWorldChannel := make ( chan *Customer )
	shopToBarberChannel := make ( chan *Customer , numberOfSeats + 1 )
	barberToShopChannel := make ( chan *Customer )
	terminationChannel := make ( chan *Customer ) 

	go barber ( hairTrimTime , shopToBarberChannel, barberToShopChannel )
	go shop ( numberOfSeats , worldToShopChannel , shopToBarberChannel , barberToShopChannel , shopToWorldChannel )
	go world ( shopToWorldChannel , terminationChannel )

	for i := 0 ; i < numberOfCustomers ; i++ {
		time.Sleep ( nextCustomerWaitTime ( ) )
		fmt.Printf ( "Customer %d enters the shop.\n" , i )
		worldToShopChannel <- & Customer { i , false }
	}

	worldToShopChannel <- & Customer { -1 , false }
	<- terminationChannel
}