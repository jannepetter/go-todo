package handlers

import (
	"context"
	"goServer/db"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos(c *gin.Context) {
	coll := db.GetCollection("todos")
	filter := bson.D{}
	cursor, err := coll.Find(c, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(c)
	var results []db.Todo
	if err = cursor.All(c, &results); err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, results)

}

func AddTodo(c *gin.Context) {
	var requestBody db.Todo
	if err := db.GetBody(c, &requestBody); err != nil {
		return
	}
	coll := db.GetCollection("todos")
	res, err := coll.InsertOne(c, requestBody)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oid, _ := res.InsertedID.(primitive.ObjectID)
	requestBody.ID = oid.Hex()
	c.JSON(http.StatusOK, requestBody)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var requestBody db.Todo
	if err := db.GetBody(c, &requestBody); err != nil {
		return
	}

	filter := bson.M{"_id": objectID}
	coll := db.GetCollection("todos")
	update := bson.M{"$set": bson.M{"title": requestBody.Title, "completed": requestBody.Completed}}

	result, err := coll.UpdateOne(c, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	if result.ModifiedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Document was not updated"})
	} else {
		requestBody.ID = objectID.Hex()
		c.JSON(http.StatusOK, requestBody)
	}

}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	filter := bson.M{"_id": objectID}

	coll := db.GetCollection("todos")
	deleteResult, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete data from MongoDB"})
		return
	}

	if deleteResult.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func SetTodoRoutes(router *gin.RouterGroup) {
	router.GET("/todos", GetTodos)
	router.POST("/todos", AddTodo)
	router.DELETE("/todos/:id", DeleteTodo)
	router.PUT("/todos/:id", UpdateTodo)
}
