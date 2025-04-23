package players

import "time"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePlayer(hashedID, encryptedID string) (*Player, error) {
	player := &Player{
		HashedPlayerID:    hashedID,
		EncryptedPlayerID: encryptedID,
		CreatedAt:         time.Now().Unix(),
	}
	if err := s.repo.Create(player); err != nil {
		return nil, err
	}
	return player, nil
}

func (s *Service) GetPlayerByHashedID(hashedID string) (*Player, error) {
	return s.repo.FindByHashedID(hashedID)
}
