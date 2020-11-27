package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Post struct {
	//ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Body  string `json:"body" bson:"body"`
}

func InsertPost(post Post) {

	collection := ConfigDB().Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	insertResult, err := collection.InsertOne(ctx, post)
	if err != nil {
		fmt.Println("Unable to insert")
	}

	fmt.Println("Inserted a post document: ", insertResult.InsertedID)
}

func GetAllPosts() {
	collection := ConfigDB().Collection("posts")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Unable to find")
	}

	var results []Post
	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents: %+v\n", results)
}
