package model

type User struct {
	Id   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
}
