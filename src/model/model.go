package model

type User struct {
	Id           string `json:"id,omitempty" bson:"_id,omitempty"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	SecondName   string `json:"second_name"`
	PasswordHash string `json:"password_hash"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
