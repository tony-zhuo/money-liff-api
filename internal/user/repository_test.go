package user

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo Repository
	user *entity.User
}

func (s *UserRepositoryTestSuite) SetupSuite() {
	db, mock, err := sqlmock.New()
	if err != nil {
		s.T().Fatalf("[UserRepositoryTestSuite](SetupTest) sqlmock error: %v", err)
	}

	dialector := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       db,
	})
	gormMockDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		s.T().Fatalf("[UserRepositoryTestSuite](SetupTest) gorm open error: %v", err)
	}

	s.mock = mock
	s.db = gormMockDB
	s.repo = NewRepository(gormMockDB, log.TeeDefault())
}

func (s *UserRepositoryTestSuite) TestRepository_Get() {
	id := 1
	lineId := "U060d21d2aedb6afeee372d9aba70b6361"
	name := "money-liff-user-get"

	rows := s.mock.NewRows([]string{"id", "line_id", "name"}).AddRow(id, lineId, name)
	s.mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE line_id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(lineId).
		WillReturnRows(rows)

	resUser := s.repo.Get(lineId)
	s.Assert().Equal(entity.User{
		Id:     id,
		LineId: lineId,
		Name:   name,
	}, *resUser)
}

func (s *UserRepositoryTestSuite) TestRepository_Create() {
	var (
		lineId = "U060d21d2aedb6afeee372d9aba70b6362"
		name   = "money-liff-user-create"
	)

	rows := s.mock.NewRows([]string{"line_id", "name"}).AddRow(lineId, name)

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("line_id","name","avatar_url","created_at","updated_at","deleted_at") VALUES ($1, $2, $3, $4, $5, $6)`)).
		WithArgs(lineId, name, "", "", "", "").
		WillReturnRows(rows)
	s.mock.ExpectCommit()

	if err := s.db.Create(&entity.User{LineId: lineId, Name: name}); err != nil {
		s.T().Errorf("[UserRepositoryTestSuite](TestRepository_Create) create data err: %+v", err)
	}

	//if err := s.mock.ExpectationsWereMet(); err != nil {
	//	s.T().Errorf("there were unfulfilled expectations: %s", err)
	//}
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
