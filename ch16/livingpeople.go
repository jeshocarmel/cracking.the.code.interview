package ch16

import (
	"math"
	"sort"
)

//Person denotes a person in the living people problem (16.10)
type Person struct {
	birth, death int
}

/* solution requires O(PY+R) time complexity where P is the number of people,
* Y is the years they are alive and R is the size of the array
 */
func maxLivingPeopleYearBF(persons []Person) int {

	var histogram = make([]int, 101)

	for _, p := range persons {
		for i := p.birth - 1900; i <= p.death-1900; i++ {
			histogram[i]++
		}
	}

	maximum, maxyear := math.MinInt8, math.MinInt8

	for i := range histogram {
		if histogram[i] > maximum {
			maximum, maxyear = histogram[i], i+1900
		}
	}

	return maxyear
}

func maxLivingPeopleOptimal(persons []Person) int {

	births := []int{}
	deaths := []int{}

	for _, p := range persons {
		births = append(births, p.birth)
		deaths = append(deaths, p.death)
	}

	sort.Slice(births, func(i, j int) bool {
		return births[i] < births[j]
	})

	sort.Slice(deaths, func(i, j int) bool {
		return deaths[i] < deaths[j]
	})

	birthIndex, deathIndex := 0, 0
	currentlyAlive := 0
	maxAlive, maxAliveYear := 0, 1989

	for birthIndex < len(births) {

		if births[birthIndex] <= deaths[deathIndex] {

			currentlyAlive++

			if currentlyAlive > maxAlive {
				maxAlive = currentlyAlive
				maxAliveYear = births[birthIndex]
			}
			birthIndex++
		} else if births[birthIndex] > deaths[deathIndex] {
			currentlyAlive--
			deathIndex++
		}
	}

	return maxAliveYear

}
