package main

type (
	Degrees float64
	Radians float64

	SchemaType struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	}

	CitySchema struct {
		Country SchemaType `json:"country"`
		City    SchemaType `json:"city"`
		Lat     SchemaType `json:"lat"`
		Lon     SchemaType `json:"lon"`
	}

	ResultSchema struct {
		Items []CitySchema `json:"items"`
	}

	City struct {
		Country string  `json:"country"`
		City    string  `json:"city"`
		Lat     Degrees `json:"lat"`
		Lon     Degrees `json:"lon"`
	}

	Cities struct {
		Items []City `json:"items"`
	}
)
