package db

import (
	"context"
	"time"

	"github.com/iduslab/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddBox(title, description string) error {
	_, err := db.Collection("box").InsertOne(context.TODO(), models.Box{
		Title:       title,
		Description: description,
		Timestamp:   time.Now(),
	})

	return err
}

func GetAllBox() (*[]models.Box, error) {
	var data []models.Box
	cursor, err := db.Collection("box").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetBox(title string) (data *models.Box, err error) {
	filter := models.Box{Title: title}
	err = db.Collection("box").FindOne(context.TODO(), filter).Decode(&data)
	return
}

func AddMemo(box primitive.ObjectID, userid string, text string) error {
	_, err := db.Collection("note").InsertOne(context.TODO(), models.Note{
		Author:    userid,
		Box:       box,
		Text:      text,
		Timestamp: time.Now(),
	})
	return err
}

func PickMemo(box primitive.ObjectID, count int) (*[]models.Note, error) {
	var data []models.Note
	memo := db.Collection("note")
	id, _ := primitive.ObjectIDFromHex(box.Hex())
	pipeline := bson.D{primitive.E{Key: "$match", Value: bson.D{primitive.E{Key: "box", Value: id}}}}
	pipeline2 := bson.D{primitive.E{Key: "$sample", Value: bson.D{primitive.E{Key: "size", Value: count}}}}
	cursor, err := memo.Aggregate(context.TODO(), mongo.Pipeline{pipeline, pipeline2})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	return &data, nil
}
