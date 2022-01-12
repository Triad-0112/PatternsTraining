package main

import "fmt"

//(1) PLAYER
type player struct {
	dress      dress
	playerType string // Counter Terrorist or Terrorist
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		playerType: playerType,
		dress:      dress,
	}
}
func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

//(2)
const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

type dressFactory struct {
	dressMap map[string]dress
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}

// FLYWEIGHT!!!
type dress interface {
	getColor() string
}

// TERRORIST DRESS
type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

// COUNTER TERRORIST DRESS
type counterTerroristDress struct {
	color string
}

func (t *counterTerroristDress) getColor() string {
	return t.color
}
func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "blue"}
}

//GAME
type game struct {
	terrorist        []*player
	counterTerrorist []*player
}

func newGame() *game {
	return &game{
		terrorist:        make([]*player, 1),
		counterTerrorist: make([]*player, 1),
	}
}
func (c *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorist = append(c.terrorist, player)
	return
}
func (c *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorist = append(c.counterTerrorist, player)
	return
}

func main() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
