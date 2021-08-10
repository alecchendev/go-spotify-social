package main

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/alecchendev/go-spotify-social/db"
)

func TestDB(t *testing.T) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	dbUri := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	client, ctx := db.InitializeDBClient(dbUri, dbName, collectionName)
	defer client.Disconnect(ctx)

	// Hopefully no one with the username testUser or testUser2 signs up...
	db.InsertUser("testUser", "Alec Chen", "i6hd3")
	db.InsertUser("testUser2", "Alec Chen2", "i6hd4")
	db.UpdateFollowing("testUser", "testUser2")
	following := db.GetFollowing("testUser2")
	if len(following) != 1 || following[0] != "testUser" {
		t.Errorf("GetFollowing returned %v (wrong output).", following)
	}
	refreshToken := db.GetRefreshToken("testUser")
	if refreshToken != "i6hd3" {
		t.Errorf("GetRefreshToken returned %v instead of \"i6hd3\".", refreshToken)
	}
	displayName := db.GetDisplayName("testUser")
	if displayName != "Alec Chen" {
		t.Errorf("GetDisplayName returned %v instead of \"Alec Chen\"", displayName)
	}
	db.DeleteUser("testUser")
	db.DeleteUser("testUser2")
}

// func TestUpdateFollowing(t *testing.T) {

// }

// func TestGetFollowing(t *testing.T) {

// }

// func TestGetRefreshToken(t *testing.T) {

// }

// func TestGetDisplayName(t *testing.T) {

// }
