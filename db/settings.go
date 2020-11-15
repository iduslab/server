package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/iduslab/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddSetting(name, description, value string) error {
	_, err := db.Collection("setting").InsertOne(context.TODO(), models.Setting{
		Name:        name,
		Description: description,
		Value:       value,
	})

	return err
}

func UpdateSettingValue(name, value string) error {
	_, err := db.Collection("setting").UpdateOne(context.TODO(), bson.D{primitive.E{Key: "name", Value: name}}, bson.D{primitive.E{Key: "value", Value: value}})
	return err
}

func GetAllSetting() (*[]models.Setting, error) {
	var data []models.Setting
	projection := bson.D{primitive.E{Key: "name", Value: 1}, primitive.E{Key: "description", Value: 1}}
	cursor, err := db.Collection("setting").Find(context.TODO(), bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetSetting(name string) (string, error) {
	var data models.Setting
	filter := models.Setting{Name: name}
	err := db.Collection("setting").FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		return "", err
	}
	return data.Value, err
}
