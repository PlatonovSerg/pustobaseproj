package players

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	HashedPlayerID    string `gorm:"size:255;not null;uniqueIndex"`
	EncryptedPlayerID string `gorm:"size:255;not null;uniqueIndex"`
	CreatedAt         int64  `gorm:"autoCreateTime"`
}
