package main

import "fmt"

type train interface {
	arrive()
	depart()
	permitArrival()
}

//PASSENGER TRAIN
type passengerTrain struct {
	mediator mediator
}

func (g *passengerTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("Passenger Train : Arrival blocked, waiting")
		return
	}
	fmt.Println("Passenger Train: Arrived")
}

func (g *passengerTrain) depart() {
	fmt.Println("Passenger Train: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *passengerTrain) permitArrival() {
	fmt.Println("Passenger Train: Arrival permitted, arriving")
	g.arrive()
}

//FREIGHT TRAIN
type freightTrain struct {
	mediator mediator
}

func (g *freightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("Freight Train: Arrival blocked, waiting")
		return
	}
	fmt.Println("Freight Train: Arrived")
}
func (g *freightTrain) depart() {
	fmt.Println("Freight Train: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *freightTrain) permitArrival() {
	fmt.Println("Freight Train: Arrival permitted")
	g.arrive()
}

//MEDIATOR REAL DEAL!!
type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}

//Station Manager
type stationManager struct {
	isPlatformFree bool
	trainQueue     []train
}

func newStationManager() *stationManager {
	return &stationManager{
		isPlatformFree: true,
	}
}

func (s *stationManager) canArrive(t train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}
func (s *stationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

func main() {
	stationManager := newStationManager()

	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}
	freightTrain := &freightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
