package structure

import (
	"time"
)

type Config struct {
	Token   string `json:"token"`
	OwnerID string `json:"ownerID"`
	Prefix  string `json:"prefix"`
}

type Help struct {
	Command     string `json:"command"`
	Usage       string `json:"usage"`
	Description string `json:"description"`
}

type Note struct {
	BoxNum    int    `gorm:"not null;"`
	Author    string `gorm:"not null;type:varchar(30)"`
	Text      string `gorm:"type:varchar(100);not null"`
	Anon      bool
	Timestamp time.Time `gorm:"not null"`
}

type Box struct {
	ID        int       `gorm:"not null"`
	Text      string    `gorm:"type:varchar(50);"`
	Timestamp time.Time `gorm:"not null"`
}

type DiscordUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   string `json:"public_flags"`
}
