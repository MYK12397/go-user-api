package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	ds "github.com/MYK12397/UserDS"
	godb "github.com/MYK12397/database"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "appplication/json")

	var singleuser ds.User
	json.NewDecoder(r.Body).Decode(&singleuser)

	checkvalidation := singleuser.ValidateInput(rw)

	if !checkvalidation {
		return
	}
	singleuser.CreatedOn = time.Now()
	singleuser.UpdateOn = time.Now()
	singleuser.Active = true
	collection := godb.ConnectDB()

	ans, err := collection.InsertOne(context.TODO(), singleuser)

	if err != nil {
		godb.GetError(err, rw)
		return
	}
	json.NewEncoder(rw).Encode(ans)
}

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	collection := godb.ConnectDB()
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		http.Error(rw, "Unable to find Id", http.StatusNotFound)
		return
	}

	defer cursor.Close(context.TODO())

	var Users []ds.User
	for cursor.Next(context.TODO()) {
		var usr ds.User

		err := cursor.Decode(&usr)
		if err != nil {
			log.Fatal("Unable to decode", err)
		}
		if usr.Active {
			Users = append(Users, usr)
		}

	}
	// if err := cursor.Err(); err != nil {
	// 	log.Fatal(err)}

	json.NewEncoder(rw).Encode(Users)
}

func GetID(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")
	var user ds.User
	vars := mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(vars["id"])
	// if er != nil {
	// 	rw.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	sorted := bson.M{"_id": id}

	collection := godb.ConnectDB()
	err := collection.FindOne(context.TODO(), sorted).Decode(&user)

	if err != nil {
		godb.GetError(err, rw)
		return
	}
	if !user.Active {
		http.Error(rw, "user does not exist", 404)
		return
	}

	json.NewEncoder(rw).Encode(user)
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal("error in ObjectIDFromHex(params) ", err)
	}
	var user ds.User
	json.NewDecoder(r.Body).Decode(&user)

	checkvalidation := user.ValidateInput(rw)

	if !checkvalidation {
		return
	}

	update := bson.D{
		{"$set", bson.D{
			{"updatedon", time.Now()},
			{"firstname", user.FirstName},
			{"lastname", user.LastName},
			{"age", bson.D{
				{"value", user.Age.Value},
				{"interval", user.Age.Interval},
			}},
			{"mobile", user.Mobile},
		}},
	}
	collection := godb.ConnectDB()
	// create filter document
	filter := bson.M{"_id": id}

	err2 := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)

	if err2 != nil {
		godb.GetError(err, rw)
		return
	}

	json.NewEncoder(rw).Encode(user)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(vars["id"])

	if err != nil {
		log.Fatal("error ino bjectIDFromHex method", err)
	}

	collection := godb.ConnectDB()
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"active": false}}

	ans, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		godb.GetError(err, rw)
		return
	}

	json.NewEncoder(rw).Encode(ans)
}
