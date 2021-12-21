package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"squidGameGo/model"
	"squidGameGo/shared"
	"time"
)

func BulkInsert(client *mongo.Client, ctx *context.Context, personList []model.SquidPlayer) error {
	persons := []interface{}{}

	for _, person := range personList {
		persons = append(persons, person)
	}
	collection := client.Database("squid").Collection("players")
	_, err := collection.InsertMany(*ctx, persons)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func MongoOpen() (*mongo.Client, *context.Context, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(shared.Config.MONGOURL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client, &ctx, err
}
