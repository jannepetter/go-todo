package db

import (
	"context"
)

func ClearTestDB() {
	Client.Database("testdata").Drop(context.TODO())
}
func BoolPointer(b bool) *bool {
	return &b
}
func InitTestData() {
	ClearTestDB()
	database := Client.Database("testdata")
	coll := database.Collection("todos")
	todos := []interface{}{
		Todo{Title: "test title", Completed: BoolPointer(true)},
		Todo{Title: "something", Completed: BoolPointer(false)},
		Todo{Title: "Another", Completed: BoolPointer(false)},
		Todo{Title: "Hello", Completed: BoolPointer(true)},
	}
	coll.InsertMany(context.TODO(), todos)
}
