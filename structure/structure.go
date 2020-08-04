package structure

import "time"

// Structure of config.json
type Config struct {
	Token   string `json:"token"`
	OwnerID string `json:"ownerID"`
	Prefix  string `json:"prefix"`
}

// Structure of IdeaNote
type Note struct {
	BoxNum    int    `gorm:"not null;"`
	Author    string `gorm:"not null;type:varchar(30)"`
	Text      string `gorm:"type:varchar(100);not null"`
	Anon      bool
	Timestamp time.Time `gorm:"not null"`
}

// Structure of Idea Box
type Box struct {
	ID        int       `gorm:"not null"`
	Text      string    `gorm:"type:varchar(50);"`
	Timestamp time.Time `gorm:"not null"`
}
