package main

import "com.safari.design_pattern/command/restaurant_example"

func main() {
	r := restaurant_example.NewRestaurant()

	tasks := []restaurant_example.Command{
		r.MakePizza(1),
		r.MakeSalad(2),
		r.CleanDishes(),
		r.MakeSalad(1),
		r.MakePizza(3),
		r.MakePizza(4),
		r.CleanDishes(),
	}

	cooks := []*restaurant_example.Cook{{}, {}}

	for i, task := range tasks {
		c := cooks[i%len(cooks)]
		c.Commands = append(c.Commands, task)
	}

	for _, c := range cooks {
		c.ExecuteCommands()
	}
}
