package users

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"

	"commercial-propfloor-users/models"

	"github.com/go-playground/validator/v10"

	"regexp"
)

func AddUserdetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		emailid := c.PostForm("emailid")
		password := c.PostForm("password")
		Userlogin(username, emailid, password)
	}
}

func Userlogin(username string, emailid string, password string) (output string) {
	godotenv.Load()
	docc := models.User{Username: username, Emailid: emailid, Password: password}
	pattern := regexp.MustCompile("^[a-zA-Z]*$")
	validate := validator.New()
	err := validate.Struct(docc)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	} else if pattern.MatchString(docc.Username) {
		usercount := Checkuserdetails(username, emailid, password)
		fmt.Println("usercount is---------------------  ", usercount)
		if usercount == 0 {
			fmt.Println("!!!!!! user does not exist")
		} else {
			fmt.Println(" LOGINED SUCCESSFULLY")
		}
	} else {
		fmt.Println("invalid input")
	}
	fmt.Println()

	return
}

func DBconnect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	fmt.Println("connected successfully")
	defer cancel()
	return client, ctx, cancel, err
}

func Checkuserdetails(user_name string, email_id string, pass_word string) (counts int64) {
	client, ctx, _, _ := DBconnect(os.Getenv("DB_HOST"))
	collection := client.Database("logindatabase").Collection("users")
	res, error := collection.CountDocuments(context.Background(), bson.M{"username": user_name, "emailid": email_id, "password": pass_word})
	if error != nil {
		fmt.Println(error)
	}
	CloseDBConnection(client, ctx)
	return res
}

func CloseDBConnection(client *mongo.Client, ctx context.Context) {
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection closed successfully.")
}
