package main

func main() {

}

func fuelConsumption(liter, consumption int) int {
	const perKm = 100
	distances := 0
	const safeDistance = 5
	distances = (perKm - safeDistance) * liter / consumption
	return distances
}
