package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MONGOURL = "mongodb://local_mongodb:27017/"
var Client *mongo.Client
var Validate *validator.Validate

func InitDb() {
	Validate = validator.New()
	clientOptions := options.Client().ApplyURI(MONGOURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn_client, err := mongo.Connect(ctx, clientOptions)
	Client = conn_client

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = Client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!", os.Getenv("ENV"))
}

func GetCollection(collection string) *mongo.Collection {
	use_db := "testdata"
	if os.Getenv("ENV") != "test" {
		use_db = "development"
	}
	database := Client.Database(use_db)
	return database.Collection(collection)
}

func GetBody(c *gin.Context, Struct any) error {
	if err := c.BindJSON(Struct); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	if err := Validate.Struct(Struct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return err
	}
	return nil
}
