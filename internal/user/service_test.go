package user

import (
	"database/sql"
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
	"time"
)

type UserServiceTestSuite struct {
	suite.Suite
	service Service
}

var (
	firstOrCreateOkLineId    string = "first-or-create-ok-line-id"
	firstOrCreateOkName      string = "first-or-create-ok-name"
	firstOrCreateOkAvatarUrl string = "https://first-or-create-ok-avatar-url"

	//firstOrCreateErrLineId    string = "first-or-create-err-line-id"
	//firstOrCreateErrName      string = "first-or-create-err-name"
	//firstOrCreateErrAvatarUrl string = "https://first-or-create-err-avatar-url"

	nullDeletedAt sql.NullTime = sql.NullTime{}
)

var (
	userRepoFirstOrCreateOkArgs = entity.User{
		LineId:    firstOrCreateOkLineId,
		Name:      firstOrCreateOkName,
		AvatarUrl: firstOrCreateOkAvatarUrl,
	}
	userRepoFirstOrCreateOkReturn = entity.User{
		Id:        1,
		LineId:    firstOrCreateOkLineId,
		Name:      firstOrCreateOkName,
		AvatarUrl: firstOrCreateOkAvatarUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt(nullDeletedAt),
	}
	//userRepoFirstOrCreateErrArgs = entity.User{
	//	LineId:    firstOrCreateErrLineId,
	//	Name:      firstOrCreateErrName,
	//	AvatarUrl: firstOrCreateErrAvatarUrl,
	//}
)

func (s *UserServiceTestSuite) SetupTest() {
	repo := new(MockedUserRepo)

	repo.On(
		"FirstOrCreate",
		&userRepoFirstOrCreateOkArgs,
		"line_id = ?",
		[]interface{}{userRepoFirstOrCreateOkArgs.LineId},
	).Return(&userRepoFirstOrCreateOkReturn, nil)

	//repo.On("FirstOrCreate",
	//	&userRepoFirstOrCreateErrArgs,
	//	"line-id = ?",
	//	userRepoFirstOrCreateErrArgs.LineId,
	//).Return(&userRepoFirstOrCreateOkReturn, nil)

	s.service = NewService(repo, log.TeeDefault())
}

func (s *UserServiceTestSuite) TestRegisterOrFindOk() {
	user, err := s.service.RegisterOrFind(&userRepoFirstOrCreateOkArgs)
	assert.Equal(s.T(), err, nil)
	assert.Equal(s.T(), *user, userRepoFirstOrCreateOkReturn)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

type MockedUserRepo struct {
	mock.Mock
}

func (m *MockedUserRepo) Get(where string, args ...interface{}) (*entity.User, error) {
	calledArgs := m.Called(where, args)
	return calledArgs.Get(0).(*entity.User), calledArgs.Error(1)
}

func (m *MockedUserRepo) Create(user *entity.User) error {
	calledArgs := m.Called(user)
	return calledArgs.Error(0)
}

func (m *MockedUserRepo) FirstOrCreate(user *entity.User, where string, args ...interface{}) (*entity.User, error) {
	calledArgs := m.Called(user, where, args)
	return calledArgs.Get(0).(*entity.User), calledArgs.Error(1)
}

func (m *MockedUserRepo) Update(user *entity.User, where string, args ...interface{}) error {
	calledArgs := m.Called(user, where, args)
	return calledArgs.Error(0)
}
