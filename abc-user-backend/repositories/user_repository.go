package repositories

import "abc-user-backend/models"

type UserRepository interface {
    GetAll(search string, limit, offset int) ([]models.User, error)
    GetByID(id string) (*models.User, error)
    Create(user *models.User) error
    Update(user *models.User) error
    Delete(id string) error
    IsEmailUnique(email string, excludeID string) (bool, error)
    Count(search string) (int64, error)
}

