package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/anika308142/mongoapi/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost/"
const dbName = "Netflix"
const colName = "Watchlist"

var collection *mongo.Collection

//connect with mongodb

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)
	//connect
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("error")
	}

	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("collection instance ready")
}

//Mongodb helpers
//inseret one record

func insertOneMovie(movie models.Netflix) *mongo.InsertOneResult {
	//var inserted *mongo.O
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie with id", inserted)
	return inserted
}

//update one record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal("Not found")
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modifified count:", result.ModifiedCount)

}

//delete one record
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal("Not found")
	}
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count:", deleteCount.DeletedCount)

}

//delete all

func deleteAllMovie() int64 {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted data total: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//read one

//read all
func GetAllMovies() []primitive.D {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.D

	for cur.Next(context.Background()) {
		var movie bson.D
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

//find one
func getOneMovie(movieId string) primitive.M {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal("Not found")
	}

	filter := bson.M{"_id": id}

	var result bson.M
	if err = collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return result

}

// controller
func AddMovie(context *gin.Context) {
	var newMovie models.Netflix
	if err := context.BindJSON(&newMovie); err != nil {
		return
	}
	result := insertOneMovie(newMovie)
	context.IndentedJSON(http.StatusCreated, result)
}

//
func MarkAsWatched(context *gin.Context) {
	id := context.Param("id")
	updateOneMovie(id)
	context.IndentedJSON(http.StatusOK, "updated")
}
func RemoveOneMovie(context *gin.Context) {
	id := context.Param("id")
	deleteOneMovie(id)
	context.IndentedJSON(http.StatusOK, "deleted")
}
func removeAllMovie(context *gin.Context) {

	count := deleteAllMovie()
	context.IndentedJSON(http.StatusOK, count)
}

func ReadAllMovie(context *gin.Context) {
	movies := GetAllMovies()
	context.IndentedJSON(http.StatusOK, movies)
}
func ReadOneMovie(context *gin.Context) {
	id := context.Param("id")
	movies := getOneMovie(id)
	context.IndentedJSON(http.StatusOK, movies)
}
