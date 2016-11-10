package main

import (
	planet "Training_GoLanguage/Solar_System/planet"
	ifPlanet "Training_GoLanguage/Solar_System/planet/interfacePlanet"
	"fmt"
)

const (
	DistanceFromSun   = 50000
	DistanceBwPlanets = 5000
)

func main() {
	mercury := planet.Planet{"Mercury", 40, 1}
	earth := planet.Planet{"Earth", 100, 3}
	jupiter := planet.Planet{"Jupiter", 5000, 5}
	pluto := planet.Planet{"Pluto", 30, 9}

	var planetList []planet.Planet
	planetList = append(planetList, mercury)
	planetList = append(planetList, earth)
	planetList = append(planetList, jupiter)
	planetList = append(planetList, pluto)

	for _, name := range planetList {
		printString(name)
		distance := name.DistanceFromSun(DistanceFromSun, DistanceBwPlanets)
		flag := isInnerPlanet(distance)
		if flag {
			fmt.Printf("The Planet %s (Inner Planet) is %d kms away from Sun\n\n", name.Name(), distance)
		} else {
			fmt.Printf("The Planet %s (Outer Planet) is %d kms away from Sun\n\n", name.Name(), distance)
		}
	}
}

func printString(planet ifPlanet.Planet) string {
	fmt.Printf("The Planet '%s'\n", planet)
	return fmt.Sprintf("The Planet '%s'", planet)
}

func isInnerPlanet(distance int) bool {
	if distance <= 70000 {
		return true
	} else {
		return false
	}
}
