package repository_test

import (
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"user-management/internal/model"
	"user-management/internal/repository"
	"user-management/internal/util"
)

type Any struct{}

func (a Any) Match(v driver.Value) bool {
	return true
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

type UserTestSuite struct {
	suite.Suite
	assert *require.Assertions
	db     *sql.DB
	dbMock sqlmock.Sqlmock
	repo   *repository.UserRepository
}

func (suite *UserTestSuite) SetupSuite() {
	suite.assert = suite.Require()

	var err error
	suite.db, suite.dbMock, err = sqlmock.New()
	suite.assert.NoError(err)

	db := sqlx.NewDb(suite.db, "mysql")
	suite.repo = repository.NewUserRepository(db)
	suite.assert.NoError(suite.dbMock.ExpectationsWereMet())
}

func (suite *UserTestSuite) Test_Store_User() {
	FirstName := "A"
	LastName := "B"
	Nickname := "C"
	Password := "1234"
	Email := "abcd@efg.com"
	Country := "UK"

	User := model.User{
		FirstName: FirstName,
		LastName:  LastName,
		Nickname:  Nickname,
		Password:  Password,
		Email:     Email,
		Country:   Country}

	suite.dbMock.ExpectExec("INSERT INTO users (.+) VALUES (.+)").
		WithArgs(
			Any{},
			User.FirstName,
			User.LastName,
			User.Nickname,
			User.Password,
			User.Email,
			User.Country,
		).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := suite.repo.Save(User)

	suite.assert.NoError(err)

	ra, err := result.RowsAffected()
	suite.assert.NoError(err)
	suite.assert.Equal(int64(1), ra)

	suite.assert.NoError(suite.dbMock.ExpectationsWereMet())
}

func (suite *UserTestSuite) Test_Update_User() {
	Id := util.GenerateUUId()
	FirstName := "A"
	LastName := "B"
	Nickname := "C"
	Password := "1234"
	Email := "abcd@efg.com"
	Country := "UK"

	User := model.User{
		Id:        Id,
		FirstName: FirstName + "A",
		LastName:  LastName,
		Nickname:  Nickname,
		Password:  Password,
		Email:     Email,
		Country:   Country}

	sqlmock.NewRows([]string{"id", "first_name", "last_name", "nickname", "password", "email", "country"}).
		AddRow(Id, FirstName, LastName, Nickname, Password, Email, Country)

	suite.dbMock.ExpectExec("UPDATE users SET (.+) WHERE (.+)").
		WithArgs(
			FirstName+"A",
			LastName,
			Nickname,
			Password,
			Email,
			Country,
			Id,
		).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := suite.repo.Update(User.Id, User)

	suite.assert.NoError(err)

	ra, err := result.RowsAffected()
	suite.assert.NoError(err)
	suite.assert.Equal(int64(1), ra)

	suite.assert.NoError(suite.dbMock.ExpectationsWereMet())
}

func (suite *UserTestSuite) Test_Delete_User() {
	Id := util.GenerateUUId()
	FirstName := "A"
	LastName := "B"
	Nickname := "C"
	Password := "1234"
	Email := "abcd@efg.com"
	Country := "UK"

	_ = sqlmock.NewRows([]string{"id", "first_name", "last_name", "nickname", "password", "email", "country"}).
		AddRow(Id, FirstName, LastName, Nickname, Password, Email, Country)

	suite.dbMock.ExpectExec("DELETE FROM users WHERE (.+)").
		WithArgs(
			Id,
		).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := suite.repo.Delete(Id)

	suite.assert.NoError(err)

	ra, err := result.RowsAffected()
	suite.assert.NoError(err)
	suite.assert.Equal(int64(1), ra)

	suite.assert.NoError(suite.dbMock.ExpectationsWereMet())
}

func (suite *UserTestSuite) Test_Get_User() {
	Id := util.GenerateUUId()
	FirstName := "A"
	LastName := "B"
	Nickname := "C"
	Password := "1234"
	Email := "abcd@efg.com"
	Country := "UK"

	row := sqlmock.NewRows([]string{"id", "first_name", "last_name", "nickname", "password", "email", "country"}).
		AddRow(Id, FirstName, LastName, Nickname, Password, Email, Country)

	suite.dbMock.ExpectQuery("SELECT .+ FROM users WHERE id .+").
		WithArgs(Id).
		WillReturnRows(row)

	user, err := suite.repo.Get(Id)
	suite.assert.NoError(err)
	suite.assert.Equal(user.Id, Id)
	suite.assert.NoError(suite.dbMock.ExpectationsWereMet())
}

func (suite *UserTestSuite) Test_Get_All_User() {
	Id1 := util.GenerateUUId()
	Id2 := util.GenerateUUId()
	FirstName := "A"
	LastName := "B"
	Nickname := "C"
	Password := "1234"
	Email := "abcd@efg.com"
	Country := "UK"

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "nickname", "password", "email", "country"}).
		AddRow(Id1, FirstName, LastName, Nickname, Password, Email, Country).
		AddRow(Id2, FirstName, LastName, Nickname, Password, Email, Country)

	suite.dbMock.ExpectQuery("SELECT .+ FROM users LIMIT .+,.+").
		WithArgs(0, 10).
		WillReturnRows(rows)

	users, err := suite.repo.GetAll(1)
	suite.assert.NoError(err)
	suite.assert.Equal(users[0].Id, Id1)
	suite.assert.Equal(users[1].Id, Id2)
	suite.assert.NoError(suite.dbMock.ExpectationsWereMet())
}

func (suite *UserTestSuite) Test_Get_User_Filter_By_Country() {
	Id := util.GenerateUUId()
	FirstName := "A"
	LastName := "B"
	Nickname := "C"
	Password := "1234"
	Email := "abcd@efg.com"
	Country := "FR"

	row := sqlmock.NewRows([]string{"id", "first_name", "last_name", "nickname", "password", "email", "country"}).
		AddRow(Id, FirstName, LastName, Nickname, Password, Email, Country)

	suite.dbMock.ExpectQuery("SELECT (.+) FROM users WHERE (.+) LIMIT (.+)").
		WithArgs(Country, 0, 10).
		WillReturnRows(row)

	users, err := suite.repo.FilterByCountry(Country, 1)
	suite.assert.NoError(err)
	suite.assert.Equal(users[0].Id, Id)
	suite.assert.NoError(suite.dbMock.ExpectationsWereMet())
}
