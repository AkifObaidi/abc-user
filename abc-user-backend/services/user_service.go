package services

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"abc-user-backend/internal/errors"
	"abc-user-backend/models"
	"abc-user-backend/repositories"
)

type UserService struct {
	repo   repositories.UserRepository
	logger *logrus.Logger
}

func NewUserService(repo repositories.UserRepository, logger *logrus.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserService) GetAll(search string, limit, offset int) ([]models.User, error) {
	return s.repo.GetAll(search, limit, offset)
}

func (s *UserService) GetByID(id string) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Create(user *models.User) error {
	if user.Age < 18 {
		return apperrors.ErrInvalidAge
	}

	unique, err := s.repo.IsEmailUnique(user.Email, "")
	if err != nil {
		return err
	}
	if !unique {
		return apperrors.ErrEmailExists
	}

	user.ID = uuid.New().String()
	return s.repo.Create(user)
}

func (s *UserService) Update(id string, user *models.User) error {
	if user.Age < 18 {
		return apperrors.ErrInvalidAge
	}

	unique, err := s.repo.IsEmailUnique(user.Email, id)
	if err != nil {
		return err
	}
	if !unique {
		return apperrors.ErrEmailExists
	}

	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	existing.Name = user.Name
	existing.Email = user.Email
	existing.Age = user.Age

	return s.repo.Update(existing)
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *UserService) Count(search string) (int64, error) {
    return s.repo.Count(search)
}
