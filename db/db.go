package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var ctx context.Context
var db *mongo.Database
var coll *mongo.Collection

func InitializeDBClient(uri, dbName, collectionName string) (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// ping
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(dbName)
	coll = db.Collection(collectionName)
	return client, ctx
}

func GetDisplayName(spotifyId string) string {

	filter := bson.D{{"spotifyId", spotifyId}}
	var result bson.M
	err := coll.FindOne(
		ctx,
		filter,
	).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			log.Fatalf("Couldn't find document with spotifyId: %v", spotifyId)
		}
		log.Fatal(err)
	}
	// fmt.Printf("found document %v", result)

	displayName, _ := result["displayName"]
	return fmt.Sprint(displayName)

	return ""
}

func GetRefreshToken(spotifyId string) string {

	filter := bson.D{{"spotifyId", spotifyId}}
	var result bson.M
	err := coll.FindOne(
		ctx,
		filter,
	).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			log.Fatalf("Couldn't find document with spotifyId: %v", spotifyId)
		}
		log.Fatal(err)
	}
	// fmt.Printf("found document %v", result)

	refreshToken, _ := result["refreshToken"]
	return fmt.Sprint(refreshToken)
}

func GetFollowing(spotifyId string) []string {

	filter := bson.D{{"spotifyId", spotifyId}}
	// cursor, err := coll.Find(ctx, filter)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var episodesFiltered []bson.M
	// if err = cursor.All(ctx, &episodesFiltered); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(episodesFiltered)

	var result bson.M
	err := coll.FindOne(
		ctx,
		filter,
	).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			log.Fatalf("Couldn't find document with spotifyId: %v", spotifyId)
		}
		log.Fatal(err)
	}
	// fmt.Printf("found document %v", result)

	v, _ := result["following"].(primitive.A)
	items := []interface{}(v)
	following := make([]string, len(items))
	for i, followingId := range items {
		following[i] = fmt.Sprint(followingId)
	}
	return following
}

func InsertUser(spotifyId, displayName, refreshToken string) {

	_, err := coll.InsertOne(ctx, bson.D{
		{"spotifyId", spotifyId},
		{"displayName", displayName},
		{"refreshToken", refreshToken},
		{"following", []string{}},
		{"followedBy", []string{}},
	})
	if err != nil {
		fmt.Println("Error inserting user.")
	}

}

func UpdateFollowing(followedId, followerId string) {

	filter := bson.D{{"spotifyId", followedId}}
	update := bson.D{{"$addToSet", bson.D{{"followedBy", followerId}}}}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("Failed to update followedBy")
	}

	filter = bson.D{{"spotifyId", followerId}}
	update = bson.D{{"$addToSet", bson.D{{"following", followedId}}}}

	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("Failed to update following")
	}
}

func DeleteUser(spotifyId string) {
	_, err := coll.DeleteMany(ctx, bson.M{"spotifyId": spotifyId})
	if err != nil {
		log.Fatal(err)
	}
}
