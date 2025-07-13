package mysql

import (
    "abc-user-backend/models"
    "abc-user-backend/repositories"
    "gorm.io/gorm"
)

type userRepo struct {
    db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repositories.UserRepository {
    return &userRepo{db: db}
}

func (r *userRepo) GetAll(search string, limit, offset int) ([]models.User, error) {
    var users []models.User
    query := r.db.Model(&models.User{})

    if search != "" {
        query = query.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
    }

    query = query.Order("created_at DESC")

    err := query.Limit(limit).Offset(offset).Find(&users).Error
    return users, err
}


func (r *userRepo) GetByID(id string) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepo) Create(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepo) Update(user *models.User) error {
    return r.db.Save(user).Error
}

func (r *userRepo) Delete(id string) error {
    return r.db.Delete(&models.User{}, "id = ?", id).Error
}

func (r *userRepo) IsEmailUnique(email string, excludeID string) (bool, error) {
    var count int64
    query := r.db.Model(&models.User{}).Where("email = ?", email)

    if excludeID != "" {
        query = query.Where("id != ?", excludeID)
    }

    err := query.Count(&count).Error
    return count == 0, err
}

func (r *userRepo) Count(search string) (int64, error) {
    var count int64
    query := r.db.Model(&models.User{})
    if search != "" {
        searchPattern := "%" + search + "%"
        query = query.Where("name LIKE ? OR email LIKE ?", searchPattern, searchPattern)
    }
    if err := query.Count(&count).Error; err != nil {
        return 0, err
    }
    return count, nil
}
