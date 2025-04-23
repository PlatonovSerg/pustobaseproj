package db

import "pustobaseproject/internal/domain/players"

func RunMigrations() error {
	return DB.AutoMigrate(&players.Player{})
}
