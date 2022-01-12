package main

import "fmt"

type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("No Brand named %s", brand)
}

//iShoe Abstraction
type iShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}
type shoe struct {
	logo string
	size int
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}
func (s *shoe) setSize(size int) {
	s.size = size
}
func (s *shoe) getLogo() string {
	return s.logo
}
func (s *shoe) getSize() int {
	return s.size
}

//iShirt Abstraction
type iShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}
type shirt struct {
	logo string
	size int
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}
func (s *shirt) getLogo() string {
	return s.logo
}
func (s *shirt) setSize(size int) {
	s.size = size
}
func (s *shirt) getSize() int {
	return s.size
}

//Concrete Factory Adidas
type adidas struct {
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}
func (a *adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

//Adidas Shoe
type adidasShoe struct {
	shoe
}

//Adidas Shirt
type adidasShirt struct {
	shirt
}

//Concrete Factory Nike
type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}
func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 14,
		},
	}
}

//Nike Shoe
type nikeShoe struct {
	shoe
}

//Nike Shirt
type nikeShirt struct {
	shirt
}

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")
	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()
	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()
	printShoeDetails(adidasShoe) 
	printShirtDetails(adidasShirt)
	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n\n", s.getSize())
}
func printShirtDetails(s iShirt) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n\n", s.getSize())
}
