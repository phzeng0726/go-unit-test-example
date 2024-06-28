package repository

import (
	"testing"

	"github.com/phzeng0726/go-unit-test-example/internal/domain"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRepositoryTestSuite struct {
	suite.Suite

	db   *gorm.DB
	repo UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.NoError(err)
	db.AutoMigrate(&domain.User{})

	suite.db = db
	suite.repo = NewUserRepository(suite.db)
}

func (suite *UserRepositoryTestSuite) TestCreateUser() {
	user := domain.User{ID: 1, Name: "Rita Zeng", Email: "rita.zeng@example.com"}

	err := suite.repo.CreateUser(user)
	suite.NoError(err)
}

func (suite *UserRepositoryTestSuite) TestCreateUser_WiteDuplicateError() {
	user := domain.User{ID: 1, Name: "Rita Zeng", Email: "rita.zeng@example.com"}

	// First user creation should succeed
	err := suite.repo.CreateUser(user)
	suite.NoError(err)

	// Second user creation with the same ID should return an error
	err = suite.repo.CreateUser(user)
	suite.Error(err)

}

func (suite *UserRepositoryTestSuite) TestCreateUser_WiteEmptyDataError() {
	emptyUser := domain.User{}

	err := suite.repo.CreateUser(emptyUser)
	suite.Error(err)
}

func (suite *UserRepositoryTestSuite) TestGetUserById() {
	suite.db.Create(&domain.User{ID: 1, Name: "Rita Zeng", Email: "rita.zeng@example.com"})

	user, err := suite.repo.GetUserById(1)
	suite.NoError(err)
	suite.NotNil(user)
	suite.Equal(1, user.ID)
	suite.Equal("Rita Zeng", user.Name)
	suite.Equal("rita.zeng@example.com", user.Email)
}

func (suite *UserRepositoryTestSuite) TestGetUserById_WithInvalidError() {
	suite.db.Create(&domain.User{ID: 1, Name: "Rita Zeng", Email: "rita.zeng@example.com"})

	_, err := suite.repo.GetUserById(2)
	suite.Error(err)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
