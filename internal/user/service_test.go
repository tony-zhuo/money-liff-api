package user

import (
	"database/sql"
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/validate_err_msg"
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
	okLineId      = "U060d21d2aedb6afeee372d9aba70b1"
	okName        = "TestName"
	okAvatarUrl   = "https://first-or-create-ok-avatar-url"
	createdAt     = time.Now()
	updatedAt     = time.Now()
	nullDeletedAt = gorm.DeletedAt(sql.NullTime{})

	lineIdValidateError = validate_err_msg.ErrorMessage{Param: "LineId", Message: "This field is required"}
	nameValidateError   = validate_err_msg.ErrorMessage{Param: "Name", Message: "This field is required"}
)

var (
	repoFirstOrCreateOkArgs = entity.User{
		LineId:    okLineId,
		Name:      okName,
		AvatarUrl: okAvatarUrl,
	}
	repoFirstOrCreateOkReturn = entity.User{
		Id:        1,
		LineId:    okLineId,
		Name:      okName,
		AvatarUrl: okAvatarUrl,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: nullDeletedAt,
	}

	repoFirstOrCreateLineIdRequiredFailedArgs = entity.User{
		LineId:    "",
		Name:      okName,
		AvatarUrl: okAvatarUrl,
	}
	repoFirstOrCreateLineIdRequiredError = validate_err_msg.ValidateErrorMessages{
		lineIdValidateError,
	}

	repoFirstOrCreateNameRequiredFailedArgs = entity.User{
		LineId:    okLineId,
		Name:      "",
		AvatarUrl: okAvatarUrl,
	}
	repoFirstOrCreateNameRequiredError = validate_err_msg.ValidateErrorMessages{
		nameValidateError,
	}

	repoGetOkArg    = "U060d21d2aedb6afeee372d9aba70b1"
	repoGetOkReturn = entity.User{
		Id:        1,
		LineId:    okLineId,
		Name:      okName,
		AvatarUrl: okAvatarUrl,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: nullDeletedAt,
	}
)

func (s *UserServiceTestSuite) SetupTest() {
	repo := new(MockedUserRepo)

	repo.On(
		"FirstOrCreate",
		&repoFirstOrCreateOkArgs,
		"line_id = ?",
		[]interface{}{repoFirstOrCreateOkArgs.LineId},
	).Return(&repoFirstOrCreateOkReturn, nil)

	repo.On("FirstOrCreate",
		&repoFirstOrCreateLineIdRequiredFailedArgs,
		"line-id = ?",
		[]interface{}{repoFirstOrCreateOkArgs.LineId},
	).Return(nil, repoFirstOrCreateLineIdRequiredError)

	repo.On("FirstOrCreate",
		&repoFirstOrCreateNameRequiredFailedArgs,
		"line-id = ?",
		[]interface{}{repoFirstOrCreateOkArgs.LineId},
	).Return(nil, repoFirstOrCreateNameRequiredError)

	repo.On("Get",
		"line_id = ?",
		[]interface{}{repoGetOkArg},
	).Return(&repoGetOkReturn, nil)

	s.service = NewService(repo, log.TeeDefault())
}

func (s *UserServiceTestSuite) TestRegisterOrFindOk() {
	user, err := s.service.RegisterOrFind(&repoFirstOrCreateOkArgs)
	assert.Equal(s.T(), nil, err)
	assert.Equal(s.T(), repoFirstOrCreateOkReturn, *user)

}

func (s *UserServiceTestSuite) TestRegisterOrFindLineIdFailed() {
	var nullUser *entity.User
	user, err := s.service.RegisterOrFind(&repoFirstOrCreateLineIdRequiredFailedArgs)
	assert.Equal(s.T(), repoFirstOrCreateLineIdRequiredError, err)
	assert.Equal(s.T(), nullUser, user)
}

func (s *UserServiceTestSuite) TestRegisterOrFindNameFailed() {
	var nullUser *entity.User
	user, err := s.service.RegisterOrFind(&repoFirstOrCreateNameRequiredFailedArgs)
	assert.Equal(s.T(), repoFirstOrCreateNameRequiredError, err)
	assert.Equal(s.T(), nullUser, user)
}

func (s *UserServiceTestSuite) TestGetOk() {
	user, err := s.service.GetUserByLineId(repoGetOkArg)
	assert.Equal(s.T(), nil, err)
	assert.Equal(s.T(), repoGetOkReturn, *user)
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
