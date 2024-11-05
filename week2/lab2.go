package main

import (
	"fmt"
)

type Snack struct {
	Name  string
	Price int
}

type Drink struct {
	Name  string
	Price int
}

func (snack *Snack) sale() {
	snack.Price = snack.Price * 9 / 10
}

func (drink *Drink) sale() {
	drink.Price = drink.Price * 8 / 10
}

func (snack *Snack) getPrice() int {
	return snack.Price
}

func (drink *Drink) getPrice() int {
	return drink.Price
}

type product interface {
	sale()
	getPrice() int
}

func SaleAndGetPrice(cart product) int {
	cart.sale()
	return cart.getPrice()
}

func main() {
	chips := Snack{"Pringles", 4000}
	cracker := Snack{"Ace", 2500}

	soda := Drink{"Sprite", 1800}
	coffee := Drink{"TOP", 2700}

	var total int = 0
	total += SaleAndGetPrice(&chips)
	total += SaleAndGetPrice(&cracker)
	total += SaleAndGetPrice(&soda)
	total += SaleAndGetPrice(&coffee)

	fmt.Println(total)
}
