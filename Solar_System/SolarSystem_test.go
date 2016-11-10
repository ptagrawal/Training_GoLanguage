package main

import (
	planet "Training_GoLanguage/Solar_System/planet"
	"testing"
)

func TestPrintString(t *testing.T) {
	testPlanet := planet.Planet{"TestPlanet", 100, 11}

	actual := printString(testPlanet)
	expected := "The Planet 'TestPlanet has mass of 100e30 Kgs'"

	if actual != expected {
		t.Errorf("Mismatching String Generated. \nExpected: %s\nGot: %s", expected, actual)
	}
}

func TestIsInnerPlanetTrue(t *testing.T) {
	distance := 50000

	res := isInnerPlanet(distance)

	if !res {
		t.Errorf("Wrong Computation done. Expected: true, Got: false")
	}
}

func TestIsInnerPlanetFalse(t *testing.T) {
	distance := 80000

	res := isInnerPlanet(distance)
	if res {
		t.Errorf("Wrong Computation done. Expected: false, Got: true")
	}
}
