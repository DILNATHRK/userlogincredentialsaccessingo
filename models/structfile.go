package models

type User struct {
	Username string `bson:"username" validate:"required"`
	Emailid  string `bson:"emailid" validate:"required"`
	Password string `bson:"password" validate:"required"`
}
