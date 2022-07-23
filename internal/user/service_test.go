package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserServiceTestSuite struct {
	suite.Suite
	service Service
}

var (
	firstOrCreateOk    = entity.User{LineId: "fake-line-id", Name: "fake-name", AvatarUrl: "https:fake-avatar"}
	firstOrCreateError = entity.User{LineId: "fake-line-id", Name: "fake-name", AvatarUrl: "https:fake-avatar"}
)

func (s *UserServiceTestSuite) SetupTest() {
	repo := new(MockedUserRepo)

	repo.On("FirstOrCreate", &firstOrCreateOk).
		Return(nil)

	s.service = NewService(repo, log.TeeDefault())
}

func (s *UserServiceTestSuite) TestCreateIfNotFound() {

}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

type MockedUserRepo struct {
	mock.Mock
}

func (m *MockedUserRepo) Get(lineId string) *entity.User {
	args := m.Called(lineId)
	return args.Get(0).(*entity.User)
}

func (m *MockedUserRepo) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockedUserRepo) FirstOrCreate(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockedUserRepo) Update(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}
