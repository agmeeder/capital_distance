package main

import "math"

func toRadians(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func haversine(city1, city2 City) float64 {
	const R = 6371

	dLat := toRadians(city2.Lat - city1.Lat)
	dLon := toRadians(city2.Lon - city1.Lon)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(toRadians(city1.Lat))*math.Cos(toRadians(city2.Lat))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
