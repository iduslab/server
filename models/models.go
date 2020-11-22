package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Config struct {
	Server struct {
		Debug bool
		Port  int
	}
	Discord struct {
		Token          string
		PermissionRole string
		Prefix         string
		ServerID       string
		ClientID       string
		ClientSecret   string
	}
	DB struct {
		Mongodb string
	}
}

type Help struct {
	Command     string `json:"command"`
	Usage       string `json:"usage"`
	Description string `json:"description"`
}

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Box       primitive.ObjectID `bson:"box,omitempty"`
	Author    string
	Text      string
	Timestamp time.Time `bson:"timestamp,omitempty"`
}

type Box struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Timestamp   time.Time          `bson:"timestamp,omitempty"`
}

type Setting struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
	Value       string             `bson:"value,omitempty" json:"value"`
}

type Role struct {
	Name   string `json:"name"`
	RoleID string `json:"role_id"`
}
