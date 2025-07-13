package services

import (
	"errors"
	"testing"

	"abc-user-backend/internal/errors"
	"abc-user-backend/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type mockLogger struct{}

func (l *mockLogger) Info(args ...interface{})  {}
func (l *mockLogger) Warn(args ...interface{})  {}
func (l *mockLogger) Error(args ...interface{}) {}
func (l *mockLogger) Fatal(args ...interface{}) {}

func newMockLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Out = nil
	return logger
}

type mockUserRepo struct {
	users               map[string]*models.User
	isEmailUniqueResult bool
	isEmailUniqueError  error
}

func (m *mockUserRepo) GetAll(search string, limit, offset int) ([]models.User, error) {
	return nil, nil
}

func (m *mockUserRepo) GetByID(id string) (*models.User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return user, nil
}

func (m *mockUserRepo) Create(user *models.User) error {
	m.users[user.ID] = user
	return nil
}

func (m *mockUserRepo) Update(user *models.User) error {
	m.users[user.ID] = user
	return nil
}

func (m *mockUserRepo) Delete(id string) error {
	delete(m.users, id)
	return nil
}

func (m *mockUserRepo) IsEmailUnique(email string, excludeID string) (bool, error) {
	return m.isEmailUniqueResult, m.isEmailUniqueError
}

func newTestUserService(emailUnique bool, emailUniqueErr error) *UserService {
	repo := &mockUserRepo{
		users:               make(map[string]*models.User),
		isEmailUniqueResult: emailUnique,
		isEmailUniqueError:  emailUniqueErr,
	}
	logger := newMockLogger()
	return NewUserService(repo, logger)
}

// -------------------- TEST CASES --------------------

func TestUserService_Create_Success(t *testing.T) {
	svc := newTestUserService(true, nil)

	user := &models.User{
		ID:    "123",
		Name:  "john",
		Email: "john@example.com",
		Age:   25,
	}

	err := svc.Create(user)
	assert.NoError(t, err)
}

func TestUserService_Create_Fails_EmailNotUnique(t *testing.T) {
	svc := newTestUserService(false, nil)

	user := &models.User{
		ID:    "123",
		Name:  "jack",
		Email: "jack@example.com",
		Age:   25,
	}

	err := svc.Create(user)
	assert.Error(t, err)
	assert.Equal(t, apperrors.ErrEmailExists, err)
}

func TestUserService_Create_Fails_InvalidAge(t *testing.T) {
	svc := newTestUserService(true, nil)

	user := &models.User{
		ID:    "124",
		Name:  "young",
		Email: "young@example.com",
		Age:   17,
	}

	err := svc.Create(user)
	assert.Error(t, err)
	assert.Equal(t, apperrors.ErrInvalidAge, err)
}