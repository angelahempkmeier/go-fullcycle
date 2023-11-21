package main

type Car struct {
	Model string
	Color string
}

func main() {
	car := Car{
		Model: "Ferrari",
		Color: "Red",
	}
	println(car.Model)
}
