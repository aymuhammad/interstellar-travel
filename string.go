package main

import "fmt"

// Create the function fuelGauge() here
func fuelGuage(fuel int) {

	fmt.Println("You have", fuel, "gallons of fuel left!")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {

	var fuel int

	if planet == "Venus" {
		fuel = 300000
	} else if planet == "Mercury" {
		fuel = 500000
	} else if planet == "Mars" {
		fuel = 700000
	} else {
		fuel = 0
	}
	return fuel

}

// Create the function greetPlanet() here
func greetPlanet(planet string) {

	fmt.Println("Welcome to planet", planet)

}

// Create the function cantFly() here
func cantFly() {

	fmt.Println("We do not have the available fuel to fly there.")

}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {

	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {

	// Test your functions!
	fmt.Println(calculateFuel("Mars"))

	// Create `planetChoice` and `fuel`
	greetPlanet("Mars")

	fmt.Println(flyToPlanet("Mars", 600000))

	fuel := 1000000
	planetChoice := "Venus"

	fuel = flyToPlanet(planetChoice, fuel)
	fuelGuage(fuel)
	// And then liftoff!

}
