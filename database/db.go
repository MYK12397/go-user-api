package database

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	fmt.Println("DataBase is Connected")

	collection := client.Database("Userdb").Collection("users_data")

	return collection
}

type ErrorResponse struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

func GetError(err error, rw http.ResponseWriter) {

	var res = ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}

	msg, _ := json.Marshal(res)
	rw.WriteHeader(res.Code)
	rw.Write(msg)
}
