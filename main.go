package main

func main() {

}

func fuelConsumption(liter, consumption int) int {
	const perKm = 100
	coveredDistances := 0
	const safeDistance = 5
	coveredDistances = (perKm - safeDistance) * liter / consumption
	return coveredDistances
}
