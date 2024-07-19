package main

import "math"

type ()

func (d Degrees) toRadians() float64 { return float64(d) * (math.Pi / 180) }

// func (r Radians) toDegrees() float64 { return float64(r) * (180.0 / math.Pi) }

func haversine(city1, city2 City) float64 {
	const R = 6371

	dLat := Degrees(city2.Lat - city1.Lat).toRadians()
	dLon := Degrees(city2.Lon - city1.Lon).toRadians()

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(city1.Lat.toRadians())*math.Cos(city2.Lat.toRadians())*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
