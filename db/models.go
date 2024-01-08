package db

type Todo struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Title     string `json:"title" validate:"required" bson:"title"`
	Completed *bool  `json:"completed" validate:"required" bson:"completed"`
}
