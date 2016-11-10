package planet

import (
	"strconv"
)

type Planet struct {
	PlanetName string
	PlanetMass int
	Position   int
}

func (planet Planet) Name() string {
	return planet.PlanetName
}

func (planet Planet) Mass() int {
	return planet.PlanetMass
}

func (planet Planet) String() string {
	return planet.Name() + " has mass of " + strconv.Itoa(planet.Mass()) + "e30 Kgs"
}

func (planet Planet) DistanceFromSun(distFromSun, distBwPlanets int) int {
	return distFromSun + distBwPlanets*planet.Position
}
