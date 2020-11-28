package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
	Body  string             `json:"body" bson:"body"`
}

func InsertPost(post Post) (interface{}, error) {

	collection := ConfigDB().Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	insertResult, err := collection.InsertOne(ctx, post)

	return insertResult.InsertedID, err
}

func GetPost(id primitive.ObjectID) (Post, error) {
	var post Post
	collection := ConfigDB().Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	return post, err
}

func DeletePost(id primitive.ObjectID) (interface{}, error) {
	collection := ConfigDB().Collection("posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return result, err
}

func UpdatePost() {

}

func GetAllPosts() ([]Post, error) {
	collection := ConfigDB().Collection("posts")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var results []Post
	for cur.Next(context.Background()) {
		//Create a value into which the single document can be decoded
		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)

	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	//Close the cursor once finished
	cur.Close(context.Background())

	return results, nil
}
