package main

import (
    "context"
	"log"
	"strings"
	schemas "github.com/atomicfruitcake/apongogo/schemas"
	readenv "github.com/atomicfruitcake/apongogo/readenv"
    "github.com/mongodb/mongo-go-driver/bson"
    "github.com/mongodb/mongo-go-driver/mongo"
    "github.com/mongodb/mongo-go-driver/mongo/options"
)


func main() {
	func connectToDB() {
		func getLoginMongoUrl(components ...string) string {
			var loginUrl string
			for _, comp := range components {
				ret += comp
			}
			return loginUrl
		}

		loginUrl := getLoginMongoUrl(
			[
				"mongodb://",readenv.mongoUser,":",readenv.mongoPass,"@",readenv.mongoUrl,":27017/"
			]
		)
	
		client, err := mongo.Connect(context.TODO(), loginUrl)
	
		if err != nil {
			log.Fatal(err)
			return err
		}
	
		err = client.Ping(context.TODO(), nil)
	
		if err != nil {
			log.Fatal(err)
			return err
		}
	
		log.Output("Successfully connected to MongoDB!")
		return client
	}

	func disconnectFromDB(client){
		err = client.Disconnect(context.TODO())

		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Output("Connection to MongoDB closed.")
	}
	
}

