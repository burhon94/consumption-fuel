package main

import "testing"

func Test_fuelConsumption(t *testing.T) {
	tests := []struct {
		name        string
		fuel        int
		consumption int
		want        int
	}{
		{name: "How much fuel is enough for kilometers", fuel: 100, consumption: 24, want: 395},
		{name: "When consumption equals existing amount", fuel: 24, consumption: 24, want: 95},
		{name: "When consumption greater than existing amount", fuel: 10, consumption: 24, want: 39},
	}
	for _, tests := range tests {
		got := fuelConsumption(tests.fuel, tests.consumption)
		if got != tests.want {
			t.Error("fuelConsumption: ", tests.name, "got: ", got, "want: ", tests.want)
		}
	}
}
