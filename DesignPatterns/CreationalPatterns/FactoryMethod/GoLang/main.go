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

type Canon struct {
	gun
}

func newCanon() iGun {
	return &Musket{
		gun: gun{
			name:  "Canon",
			power: 100,
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
	if gunType == "canon" {
		return newCanon(), nil
	}
	return nil, fmt.Errorf("Wrong gun passed")
}
func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")
	canon, _ := getGun("canon")
	printDetails(ak47)
	printDetails(musket)
	printDetails(canon)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n\n", g.getPower())
}
