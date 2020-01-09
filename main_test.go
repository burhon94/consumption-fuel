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
	}
	for _, tests := range tests {
		got := fuelConsumption(tests.fuel, tests.consumption)
		if got != tests.want {
			t.Error("fuelConsumption: ", tests.name, "got: ", got, "want: ", tests.want)
		}
	}
}
