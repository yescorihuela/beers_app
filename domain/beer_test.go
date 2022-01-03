package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	UnexpectedDatabaseError = "Unexpected database error"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository BeerRepository
	beer       Beer
	beers      []Beer
}

func getDSN() string {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	urlDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	return urlDSN
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		s.T().Errorf("Failed to open mock sql db, got error: %v", err)
	}
	if db == nil {
		s.T().Error("Mock DB is null")
	}
	if s.mock == nil {
		s.T().Error("sqlMock is null")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		Conn:                 db,
		DriverName:           "postgres",
		PreferSimpleProtocol: true,
	})
	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		s.T().Errorf("Unexpected error: %v", err)
	}
	s.repository = NewBeerRepositoryDatabase(s.DB)
	s.beer = NewBeer(1, "Torobayo", "Kunstmann", "Chile", "CLP", 1000)
	s.beers = []Beer{
		NewBeer(1, "Pale Ale", "Kunstmann", "Chile", "CLP", 1100),
		NewBeer(2, "Torobayo", "Kunstmann", "Chile", "CLP", 1000),
		NewBeer(3, "Kristal Zero", "Kristal", "Chile", "CLP", 800),
	}
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

// Not working test
// func (s *Suite) TestBeerRepositoryCreate() {

// 	beerID := uint(1)

// 	newRow := sqlmock.NewRows([]string{"id", "name", "brewery", "country", "currency", "price"}).
// 		AddRow(beerID, "Torobayo", "Kunstmann", "Chile", "CLP", 1000)

// 	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO beers(id, name, brewery, country, currency, price) VALUES($1, $2, $3, $4, $5, $6) RETURNING "beers"."id";`)).
// 		WithArgs(beerID, "Torobayo", "Kunstmann", "Chile", "CLP", 1000).
// 		WillReturnRows(newRow)

// 	_, err := s.repository.Create(s.beer)
// 	require.NoError(s.T(), errors.New(err.Message))
// }

func (s *Suite) TestBeerRepositoryFindOne() {
	beerID := 1
	row := sqlmock.NewRows([]string{"id", "name", "brewery", "country", "currency", "price"}).
		AddRow(1, "Torobayo", "Kunstmann", "Chile", "CLP", 1000)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT (.*) FROM "beers" WHERE (id = $1)`)).
		WithArgs(strconv.Itoa(beerID)).
		WillReturnRows(row)

	res, err := s.repository.FindOne(beerID)

	require.NoError(s.T(), errors.New(err.Message))
	require.Nil(s.T(), deep.Equal(&s.beer, res))
}

func (s *Suite) TestBeerRepositoryFindAll() {

	beerRows := sqlmock.NewRows([]string{"id", "name", "brewery", "country", "currency", "price"}).
		AddRow(1, "Torobayo", "Kunstmann", "Chile", "CLP", 1000).
		AddRow(2, "Pale Ale", "Kunstmann", "Chile", "CLP", 1100).
		AddRow(3, "Kristal Zero", "Kristal", "Chile", "CLP", 800)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "beers"`)).
		WillReturnRows(beerRows)

	res, err := s.repository.FindAll()
	fmt.Println(s.beers)
	require.NoError(s.T(), errors.New(err.Message))
	require.Nil(s.T(), deep.Equal(s.beers, res))
}

func (s *Suite) TestBeerRepositoryFindAllError() {

	beerRows := sqlmock.NewRows([]string{"id", "name", "brewery", "country", "currency", "price"}).
		AddRow(1, "Torobayo", "Kunstmann", "Chile", "CLP", 1000).
		AddRow(2, "Pale Ale", "Kunstmann", "Chile", "CLP", 1100).
		AddRow(3, "Kristal Zero", "Kristal", "Chile", "CLP", 800)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "beers"`)).
		WillReturnRows(beerRows)

	_, err := s.repository.FindAll()

	require.Error(s.T(), errors.New(err.Message))
}
