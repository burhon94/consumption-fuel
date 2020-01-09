package main

func main() {

}

func fuelConsumption(litr, consumption int) int {
	const per_km = 100
	distances := 0
	const safedistance = 5
	distances = (per_km - safedistance) * litr  / consumption
	return distances
}