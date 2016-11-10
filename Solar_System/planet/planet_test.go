package planet

import (
	"testing"
)

func TestString(t *testing.T) {
	planet := Planet{"TestPlanet", 100, 11}

	actual := planet.String()
	expected := "TestPlanet has mass of 100e30 Kgs"

	if actual != expected {
		t.Errorf("String Function generates incorrect response")
	}
}

func TestDistanceFromSun(t *testing.T) {
	planet := Planet{"TestPlanet", 100, 11}

	actual := planet.DistanceFromSun(10, 2)
	expected := 32

	if actual != expected {
		t.Errorf("Expected distance: %d, Got: %d", expected, actual)
	}
}
