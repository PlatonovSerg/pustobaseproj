package players

import (
	"errors"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}
func (r *GormRepository) Create(player *Player) error {
	return r.db.Create(player).Error
}

func (r *GormRepository) FindByHashedID(hash string) (*Player, error) {
	var player Player
	if err := r.db.Where("hashed_player_id = ?", hash).First(&player).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &player, nil
}
