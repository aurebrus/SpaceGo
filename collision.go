package main

import "math"

type circle struct {
	center vector
	radius float64
}

func collision(c1, c2 circle) bool {
	distance := math.Sqrt(math.Pow(c2.center.x-c1.canter.x, 2) +
		math.Pow(c2.center.y-c1.canter.y, 2))
	return distance <= c1.radius+c2.radius
}
