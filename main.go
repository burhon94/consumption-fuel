package main

func main() {

}

func fuelConsumption(liter, consumption int) int {
	const per_km = 100
	distances := 0
	const safedistance = 5
	distances = (per_km - safedistance) * liter / consumption
	return distances
}
