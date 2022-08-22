package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed rand number generator
	rand.Seed(time.Now().UnixNano())

	// opening statement
	color.Cyan("---------------------------")
	color.Cyan("The Sleeping Barber Problem")
	color.Cyan("---------------------------")

	// create channels
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create barber shop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HaircutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarbersDoneChan: doneChan,
		ClientsChan:     clientChan,
		Open:            true,
	}

	color.Green("The shop is open for the day")

	// add barbers
	shop.addBarber("Mark")
	shop.addBarber("Frank")
	shop.addBarber("Susane")
	shop.addBarber("Jan")
	shop.addBarber("Michael")
	shop.addBarber("Dwight")

	// start the barbershop go routine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShop()
		closed <- true
	}()

	// add clients
	i := 1

	go func() {
		for {
			// get a random number with average arrival rate
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is open
	<-closed

	time.Sleep(5 * time.Second)
}
