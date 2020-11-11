package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Config struct {
	Token          string `json:"token"`
	OwnerID        string `json:"ownerID"`
	Prefix         string `json:"prefix"`
	Mongodb        string `json:"mongodb"`
	PermissionRole string `json:"permissionRole"`
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

type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   string `json:"public_flags"`
}
