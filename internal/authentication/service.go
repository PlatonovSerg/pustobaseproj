package authentication

import (
	"errors"

	"pustobaseproject/pkg/utils"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

var ErrInvalidUserID = errors.New("invalid user_id")

func (s *Service) GenerateToken(userID string) (string, error) {
	if userID == "" {
		return "", ErrInvalidUserID
	}
	return utils.GenerateToken(userID)
}
