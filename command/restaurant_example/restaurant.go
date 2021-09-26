package restaurant_example

import "fmt"

type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

func NewRestaurant() *Restaurant {
	const totalDishes = 10
	return &Restaurant{
		TotalDishes:   totalDishes,
		CleanedDishes: totalDishes,
	}
}

func (r *Restaurant) MakePizza(n int) Command {
	return &MakePizzaCommand{
		n:          n,
		Restaurant: r,
	}
}

func (r *Restaurant) MakeSalad(n int) Command {
	return &MakeSaladCommand{
		n:          n,
		Restaurant: r,
	}
}

func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{
		Restaurant: r,
	}
}

type MakePizzaCommand struct {
	n int
	*Restaurant
}

func (m *MakePizzaCommand) Execute() {
	m.CleanedDishes -= m.n
	fmt.Printf("made %d pizza \n", m.n)
}

type MakeSaladCommand struct {
	n int
	*Restaurant
}

func (m *MakeSaladCommand) Execute() {
	m.CleanedDishes -= m.n
	fmt.Printf("made %d salad \n", m.n)
}

type CleanDishesCommand struct {
	*Restaurant
}

func (c CleanDishesCommand) Execute() {
	c.CleanedDishes = c.TotalDishes
	fmt.Println("dishes cleaned")
}
