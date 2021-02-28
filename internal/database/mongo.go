package database

import (
	"context"
	"fmt"
	"time"

	"github.com/Kungfucoding23/API-Go-mysql-chi/internal/logs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Mongo ...
type Mongo struct {
	Client *mongo.Client
}

//NewMongoClient ...
func NewMongoClient(host string) *Mongo {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", host)))

	if err != nil {
		logs.Log().Errorf("cannot connect with mongo %s", err.Error())
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		logs.Log().Errorf("cannot connect with mongo %s", err.Error())
		panic(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		logs.Log().Errorf("cannot PING the mongo server: %s", err.Error())
		panic(err)
	}

	return &Mongo{client}
}
