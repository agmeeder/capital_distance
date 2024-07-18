package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ollama/ollama/api"
)

const (
	model = "llama3:8b"
)

func main() {
	var country1 = "Germany"
	var country2 = "Netherlands"

	if len(os.Args) == 3 {
		country1 = os.Args[1]
		country2 = os.Args[2]
	}

	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	citySchema := CitySchema{
		Country: SchemaType{
			Type:        "string",
			Description: "The name of the country",
		},
		City: SchemaType{
			Type:        "string",
			Description: "Name of the city",
		},
		Lat: SchemaType{
			Type:        "float",
			Description: "Decimal Latitude of the city",
		},
		Lon: SchemaType{
			Type:        "float",
			Description: "Decimal Longitude of the city",
		},
	}

	resultSchema := ResultSchema{
		Items: []CitySchema{citySchema},
	}

	resultSchemaJson, err := json.Marshal(resultSchema)
	if err != nil {
		log.Fatal(err)
	}

	messages := []api.Message{
		{
			Role:    "system",
			Content: fmt.Sprintf("You are a helpful AI assistant. The user will enter two country names and the assistant will return the country, the capital per country and the decimal latitude and decimal longitude of the capital of both countries, Output in JSON using the schema defined here: %v", string(resultSchemaJson)),
		},
		{
			Role:    "user",
			Content: "Netherlands,Germany",
		},
		{
			Role:    "assistant",
			Content: "[{\"country\":\"Netherlands\", \"city\":\"Amsterdam\", \"lat\": 52.370216, \"lon\": 4.8951667 },{\"country\":\"Belgium\", \"city\":\"Brussels\", \"lat\": 50.8503395, \"lon\": 4.3516263 }]",
		},
		{
			Role:    "user",
			Content: fmt.Sprintf("%v, %v", country1, country2),
		},
	}

	stream := false
	options := make(map[string]interface{})
	options["temperature"] = 0.0
	ctx := context.Background()
	req := &api.ChatRequest{
		Model:    model,
		Messages: messages,
		Format:   "json",
		Options:  options,
		Stream:   &stream,
	}

	respFunc := func(resp api.ChatResponse) error {
		var cities Cities
		json.Unmarshal([]byte(resp.Message.Content), &cities)
		fmt.Println(cities)

		if len(cities.Items) == 2 {
			city1 := cities.Items[0]
			city2 := cities.Items[1]
			fmt.Printf("The distance between %v the capital of %v and %v the capital of %v is %.0f km.\n", city1.City, city1.Country, city2.City, city2.Country, haversine(city1, city2))
		}
		return nil
	}

	err = client.Chat(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
}
