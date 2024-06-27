package repository

import (
	"testing"

	"github.com/phzeng0726/go-unit-test-example/internal/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&domain.User{})

	return db
}

func TestCreateUser(t *testing.T) {
	db := initDatabase(t)

	user := domain.User{ID: 1, Name: "Rita Zeng", Email: "rita.zeng@example.com"}

	repo := NewUserRepository(db)

	t.Run("Create user", func(t *testing.T) {
		err := repo.CreateUser(user)
		assert.NoError(t, err)
	})

	t.Run("Create duplicate user", func(t *testing.T) {
		err := repo.CreateUser(user)
		assert.Error(t, err)
	})

	t.Run("Create user with empty data", func(t *testing.T) {
		emptyUser := domain.User{}
		err := repo.CreateUser(emptyUser)
		assert.Error(t, err)
	})
}

func TestGetUserById(t *testing.T) {
	db := initDatabase(t)
	db.Create(&domain.User{ID: 1, Name: "Rita Zeng", Email: "rita.zeng@example.com"})

	repo := NewUserRepository(db)

	t.Run("Get exist user by id", func(t *testing.T) {
		user, err := repo.GetUserById(1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "Rita Zeng", user.Name)
		assert.Equal(t, "rita.zeng@example.com", user.Email)
	})

	t.Run("Get non-existing user by id", func(t *testing.T) {
		_, err := repo.GetUserById(2)
		assert.Error(t, err)
	})
}
