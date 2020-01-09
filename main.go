package main

func main() {

}

func fuelConsumption(liter, consumption int) int {
	const perKm = 100
	distances := 0
	const safedistance = 5
	distances = (perKm - safedistance) * liter / consumption
	return distances
}
