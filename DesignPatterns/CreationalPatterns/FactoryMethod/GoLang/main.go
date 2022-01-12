package main

import "fmt"

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}
func (g *gun) getName() string {
	return g.name
}
func (g *gun) setPower(power int) {
	g.power = power
}
func (g *gun) getPower() int {
	return g.power
}

// AK 47
type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 Gun",
			power: 4,
		},
	}
}

// Musket
type Musket struct {
	gun
}

func newMusket() iGun {
	return &Musket{
		gun: gun{
			name:  "Musket",
			power: 1,
		},
	}
}
func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun passed")
}
func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}
