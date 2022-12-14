package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HaircutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (b *BarberShop) addBarber(barber string) {
	b.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			// if there are no clients, the barber goes to sleep
			if len(b.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes a nap.", barber)
				isSleeping = true
			}

			client, shopOpen := <-b.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				// cut hair
				b.cutHair(barber, client)
			} else {
				// shop is closed, sent the barber home and close go routine
				b.sendBarberHome(barber)
				return
			}
		}

	}()
}

func (b *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(b.HaircutDuration)
	color.Green("%s is finished cutting %s's hair", barber, client)
}

func (b *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	b.BarbersDoneChan <- true
}

func (b *BarberShop) closeShop() {
	color.Red("Closing shop for the day.")

	close(b.ClientsChan)
	b.Open = false

	for a := 1; a <= b.NumberOfBarbers; a++ {
		<-b.BarbersDoneChan
	}

	close(b.BarbersDoneChan)

	color.Red("--------------------------------------------------------------------")
	color.Red("The barbershop is now closed for the day and everyone has gone home.")
	color.Red("--------------------------------------------------------------------")
}

func (b *BarberShop) addClient(client string) {
	color.Green("*** %s arrives.", client)

	if b.Open {
		select {
		case b.ClientsChan <- client:
			color.Green("%s takes a seat in the waiting room.", client)
		default:
			color.Red("The waiting room is full, so %s leaves.", client)
		}
	} else {
		color.Red("The shop is already closed, so %s leaves.", client)
	}
}
