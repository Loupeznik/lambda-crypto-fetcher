package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	client "github.com/bozd4g/go-http-client"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Coin struct {
	Status struct {
		Timestamp time.Time `json:"timestamp"`
	} `json:"status"`
	Data struct {
		Symbol     string `json:"symbol"`
		Name       string `json:"name"`
		MarketData struct {
			Value float32 `json:"price_usd"`
		} `json:"market_data"`
	} `json:"data"`
}

func main() {
	lambda.Start(handle)
}

func handle() {
	assets := [3]string{"btc", "eth", "doge"}

	httpClient := client.New("https://data.messari.io/api/v1/")

	db, err := mongo.NewClient(
		options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = db.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect(ctx)

	for _, asset := range assets {
		request, err := httpClient.Get(
			"assets/" + asset + "/metrics?fields=symbol,name,market_data")

		if err != nil {
			panic(err)
		}

		response, err := httpClient.Do(request)

		if err != nil {
			panic(err)
		}

		data := Coin{}
		err = json.Unmarshal(response.Get().Body, &data)
		if err != nil {
			panic(err)
		}

		collection := db.Database("coins").Collection("coins")
		result, err := collection.InsertOne(ctx, data)

		if err != nil {
			log.Printf("Could not save values for " + data.Data.Symbol)
		} else {
			log.Printf(
				"Saved data for " + data.Data.Symbol +
					" as " + fmt.Sprint(result.InsertedID))
		}
	}
}
